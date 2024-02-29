package types

func Basic() []interface{} {
	return []interface{}{
		&Employee{},         // employee
		&LoginInformation{}, // login_information
		&Role{},             // role
		&Permission{},       // permission
		&Menu{},             // menu
		&Department{},       //department
	}
}

func M2M() []interface{} {
	return []interface{}{
		&RoleMenu{},
		&RolePermission{},
	}
}
