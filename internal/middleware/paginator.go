package middleware

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/internal/utils/paginator"
)

type PageModel struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

func Paginator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p PageModel
		// 獲取分頁資訊
		if err := gbhttp.ParseQuery(c, &p); err != nil {
			//response.FailWithMessage(c, err.Error())
			c.Abort()
			return
		}

		// 寫入上下文
		gbhttp.Set(c, "paginator", paginator.New(p.PageSize, (p.Page-1)*p.PageSize))

		c.Next()
	}
}
