package upload

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/upload/v1/api"
)

func New() gbhttp.IBind {
	return &System{}
}

type System struct{}

func (s *System) Init(group *gin.RouterGroup) {
	uploadV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(uploadV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.UploadApi{},
	}
}
