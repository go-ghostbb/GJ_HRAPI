package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	"hrapi/apps/upload/v1/model"
	"hrapi/apps/upload/v1/service"
	"hrapi/internal/middleware"
	"mime/multipart"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type UploadApi struct{}

func (u *UploadApi) Init(group *gin.RouterGroup) {
	uploadGroup := group.Group("upload", middleware.Auth())
	uploadGroup.POST("", u.upload)
}

// 上傳頭像
//
//	route => POST /api/v1/upload
func (u *UploadApi) upload(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		file *multipart.FileHeader
		out  model.UploadRes
		err  error
	)

	file, err = c.FormFile("file")
	if err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidForm, err.Error())
		return
	}

	out, err = service.Upload(ctx).Upload(file)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}
