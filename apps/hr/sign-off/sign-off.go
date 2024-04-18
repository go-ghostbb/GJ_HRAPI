package signoff

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/sign-off/v1/api"
)

func New() gbhttp.IBind {
	return &SignOff{}
}

type SignOff struct{}

func (s *SignOff) Init(group *gin.RouterGroup) {
	signoffV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(signoffV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.LeaveApi{},
		&v1.CheckInApi{},
	}
}
