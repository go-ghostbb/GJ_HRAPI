package middleware

import (
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"hrapi/internal/utils/jwtx"
	"hrapi/internal/utils/jwtx/claims"
	"hrapi/internal/utils/token"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

var (
	sf = singleflight.Group{}
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := gbhttp.GetBearerToken(c)
		if accessToken == "" {
			Responder(Mount(c)).Fail(CodeAuthTokenIllegal)
			c.Abort()
			return
		}

		// if the token is in the blacklist
		if exist, _ := token.IsBlack(g.Redis(), accessToken); exist {
			Responder(Mount(c)).Fail(CodeAuthTokenIllegal)
			c.Abort()
			return
		}

		accessClaims, err := jwtx.Service.ParseAccessToken(accessToken)
		if err != nil {
			// If the error is that the token has expired
			// and the token is not in the whitelist
			// then perform token refresh action
			if gberror.Is(err, jwtx.TokenExpired) {
				// if the token is in the whitelist, pass it
				if exist, _ := token.IsWhite(g.Redis(), accessToken); !exist {
					newAccessToken, err, _ := refresh(accessToken, accessClaims)
					if err != nil {
						Responder(Mount(c)).Fail(CodeAuthTokenExpired)
						c.Abort()
						return
					}
					if newAccessToken != nil {
						// Put the old accessToken into the whitelist
						_ = token.SetWhite(g.Redis(), accessToken)
						// Add new accessToken response header
						c.Header("x-token", newAccessToken.(string))
					}
				}
			} else {
				Responder(Mount(c)).FailWithDetail(CodeAuthTokenIllegal, err.Error())
				c.Abort()
				return
			}
		}

		// Write context
		gbhttp.Set(c, "employee.id", accessClaims.BaseClaims.EmployeeID)
		gbhttp.Set(c, "employee.real_name", accessClaims.RealName)
		gbhttp.Set(c, "employee.account", accessClaims.Account)
		gbhttp.Set(c, "employee.roles", accessClaims.Roles)
		gbhttp.Set(c, "auth.type", accessClaims.Type)
		c.Next()
	}
}

// Refresh token
func refresh(accessToken string, accessClaims *claims.AccessClaims) (interface{}, error, bool) {
	// Use singleflight to merge requests to avoid high concurrency
	return sf.Do("JWT:"+accessToken, func() (interface{}, error) {
		key := string(accessClaims.Type) + ":" + gbconv.String(accessClaims.BaseClaims.EmployeeID)

		// Get the refreshToken from redis
		refreshToken, err := token.GetRefreshToken(g.Redis(), key)
		if err != nil {
			return nil, err
		}
		// refresh
		newAccessToken, newRefreshToken, err := jwtx.Service.Refreshing(accessClaims.Type, accessToken, refreshToken)
		if err != nil {
			return nil, err
		}
		// Put the new refreshToken back into redis
		if err = token.SetRefreshToken(g.Redis(), key, newRefreshToken, jwtx.Service.GetRefreshTokenEp()); err != nil {
			return nil, err
		}
		return newAccessToken, nil
	})
}
