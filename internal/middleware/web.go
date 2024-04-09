package middleware

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/internal/types/enum"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		authType := gbhttp.Get(c, "auth.type").Interface().(enum.MenuShow)
		if authType != enum.Web {
			Responder(Mount(c)).Fail(CodeNotAuthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
