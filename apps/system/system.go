package system

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	"github.com/gin-gonic/gin"
	v1 "hrapi/apps/system/v1/api"
)

func New() gbhttp.IBind {
	return &System{}
}

type System struct{}

func (s *System) Init(group *gin.RouterGroup) {
	systemV1 := group.Group("api/v1")
	for _, r := range v1List() {
		r.Init(systemV1)
	}
}

func v1List() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.BaseApi{},
		&v1.EmployeeApi{},
		&v1.MenuApi{},
		&v1.DepartmentApi{},
		&v1.RoleApi{},
		&v1.PermissionApi{},
	}
}
