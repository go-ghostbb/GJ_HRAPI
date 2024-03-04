package middleware

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/internal/utils/paginator"
	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type PageModel struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

func Paginator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p PageModel
		// 獲取分頁資訊
		if err := gbhttp.ParseQuery(c, &p); err != nil {
			Responder(Mount(c)).FailWithDetail(CodeRequestInvalidQuery, err.Error())
			c.Abort()
			return
		}

		page := paginator.New(p.PageSize, (p.Page-1)*p.PageSize)

		// 寫入上下文
		gbhttp.Set(c, "paginator", page)

		c.Next()
	}
}
