package types

func Basic() []interface{} {
	return []interface{}{
		&Employee{},                  // employee
		&LoginInformation{},          // login_information
		&Role{},                      // role
		&Permission{},                // permission
		&Menu{},                      // menu
		&Department{},                //department
		&Leave{},                     // leave
		&LeaveGroup{},                // leave_group
		&LeaveGroupCondition{},       // leave_group_condition
		&Vacation{},                  // vacation
		&VacationGroup{},             // vacation_group
		&VacationGroupOvertimeRate{}, // vacation_group_overtime_rate
		&VacationSchedule{},          // vacation_schedule
	}
}

func M2M() []interface{} {
	return []interface{}{
		&RoleMenu{},
		&RolePermission{},
		&LeaveGroupEmployee{},
		&VacationGroupEmployee{},
	}
}
