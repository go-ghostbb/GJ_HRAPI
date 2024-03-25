package filing

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/filing/v1/api"
)

func New() gbhttp.IBind {
	return &Filing{}
}

type Filing struct{}

func (s *Filing) Init(group *gin.RouterGroup) {
	filingV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(filingV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.CheckInStatusApi{},
	}
}
