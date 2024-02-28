package middleware

import (
	gbi18n "ghostbb.io/gb/i18n/gb_i18n"
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			locale    = c.GetHeader("Locale")
			serverCtx = gbhttp.Ctx(c)
			ctx       = gbi18n.WithLanguage(serverCtx, locale)
		)
		*c.Request = *c.Request.WithContext(ctx)
		c.Next()
	}
}
