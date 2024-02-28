package jwtx

import (
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	gbutil "ghostbb.io/gb/util/gb_util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"hrapi/internal/utils/jwtx/claims"
	"time"
)

func init() {
	var (
		configMap g.Map
		err       error
	)

	if configMap, err = g.Config().Data(gbctx.New()); err != nil {
		panic(err)
	}

	if len(configMap) > 0 {
		if _, v := gbutil.MapPossibleItemByKey(configMap, "Jwt"); v != nil {
			if err = gbconv.Struct(v, Service); err != nil {
				panic(err)
			}
		}
	}
}

var (
	TokenExpired    = gberror.New("token is expired")
	RefreshTokenErr = gberror.New("wrong refresh token")
	Service         = new(Jwtx)
)

type Jwtx struct {
	SigningKey     []byte        `json:"signingKey"`
	Issuer         string        `json:"issuer"`
	AccessTokenEp  time.Duration `json:"accessTokenEp"`
	RefreshTokenEp time.Duration `json:"refreshTokenEp"`
}

func New(signingKey string, issuer string, accessTokenEp, refreshTokenEp time.Duration) *Jwtx {
	return &Jwtx{
		SigningKey:     []byte(signingKey),
		Issuer:         issuer,
		AccessTokenEp:  accessTokenEp,
		RefreshTokenEp: refreshTokenEp,
	}
}

func (j *Jwtx) GetRefreshTokenEp() time.Duration {
	return j.RefreshTokenEp
}

func (j *Jwtx) GetAccessTokenEp() time.Duration {
	return j.AccessTokenEp
}

// CreateToken 創建Token
func (j *Jwtx) CreateToken(baseClaims claims.BaseClaims) (accessToken, refreshToken string, err error) {
	key, _ := uuid.NewUUID()

	accessClaims := claims.AccessClaims{
		BaseClaims: baseClaims,
		Key:        key,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenEp)), // 過期時間，根據配置
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      // 簽發時間
			NotBefore: jwt.NewNumericDate(time.Now()),                      // 生效時間
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(j.SigningKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := claims.RefreshClaims{
		Key: key,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenEp)), // 過期時間，根據配置
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       // 簽發時間
			NotBefore: jwt.NewNumericDate(time.Now()),                       // 生效時間
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(j.SigningKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ParseAccessToken 解析AccessToken
func (j *Jwtx) ParseAccessToken(tokenString string) (*claims.AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if _claims, ok := token.Claims.(*claims.AccessClaims); ok && token.Valid {
		return _claims, nil
	} else {
		if gberror.Is(err, jwt.ErrTokenExpired) {
			return _claims, TokenExpired
		}
		return nil, err
	}
}

// ParseRefreshToken 解析refreshToken
func (j *Jwtx) ParseRefreshToken(tokenString string) (*claims.RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if _claims, ok := token.Claims.(*claims.RefreshClaims); ok && token.Valid {
		return _claims, nil
	} else {
		return nil, err
	}
}

// Refreshing 刷新token
func (j *Jwtx) Refreshing(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	accessClaims, _, ok, err := j.compareAccessAndRefresh(accessToken, refreshToken)
	if err != nil {
		return "", "", err
	}
	if !ok {
		return "", "", RefreshTokenErr
	}
	newAccessToken, newRefreshToken, err = j.CreateToken(accessClaims.BaseClaims)
	if err != nil {
		return "", "", err
	}
	return newAccessToken, newRefreshToken, nil
}

func (j *Jwtx) compareAccessAndRefresh(accessToken, refreshToken string) (*claims.AccessClaims, *claims.RefreshClaims, bool, error) {
	accessClaims, err := j.ParseAccessToken(accessToken)
	if err != nil {
		if !gberror.Is(err, TokenExpired) {
			return nil, nil, false, err
		}
	}
	refreshClaims, err := j.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, false, err
	}
	if accessClaims.Key == refreshClaims.Key {
		return accessClaims, refreshClaims, true, nil
	} else {
		return accessClaims, refreshClaims, false, nil
	}
}
