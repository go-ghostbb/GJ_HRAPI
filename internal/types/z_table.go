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
