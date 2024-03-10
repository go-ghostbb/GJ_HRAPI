package middleware

import (
	gbi18n "ghostbb.io/gb/i18n/gb_i18n"
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			locale    = c.GetHeader("x-locale")
			serverCtx = gbhttp.Ctx(c)
		)
		if locale == "" {
			locale = "en"
		}
		*c.Request = *c.Request.WithContext(gbi18n.WithLanguage(serverCtx, locale))
		c.Next()
	}
}
