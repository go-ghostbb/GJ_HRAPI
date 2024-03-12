package setting

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/setting/v1/api"
)

func New() gbhttp.IBind {
	return &Setting{}
}

type Setting struct{}

func (s *Setting) Init(group *gin.RouterGroup) {
	settingV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(settingV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.LeaveApi{},
		&v1.VacationApi{},
		&v1.WorkShiftApi{},
	}
}
