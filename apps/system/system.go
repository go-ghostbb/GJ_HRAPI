package system

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	v1 "hrapi/apps/system/api/v1"
)

func New() []gbhttp.IBind {
	return []gbhttp.IBind{
		&v1.BaseApi{},
		&v1.EmployeeApi{},
		&v1.MenuApi{},
	}
}
