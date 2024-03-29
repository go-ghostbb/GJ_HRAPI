package salary

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/hr/salary/v1/api"
)

func New() gbhttp.IBind {
	return &Salary{}
}

type Salary struct{}

func (s *Salary) Init(group *gin.RouterGroup) {
	salaryV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(salaryV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.SalaryApi{},
	}
}
