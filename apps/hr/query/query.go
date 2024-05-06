package query

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/query/v1/api"
)

func New() gbhttp.IBind {
	return &Query{}
}

type Query struct{}

func (q *Query) Init(group *gin.RouterGroup) {
	queryV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(queryV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.CheckInApi{},
	}
}
