package daily

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/daily/v1/api"
)

func New() gbhttp.IBind {
	return &Daily{}
}

type Daily struct{}

func (s *Daily) Init(group *gin.RouterGroup) {
	dailyV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(dailyV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.LeaveApi{},
		&v1.CheckInApi{},
	}
}
