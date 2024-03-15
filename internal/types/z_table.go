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
		&WorkShift{},                 // work_shift
		&WorkSchedule{},              // work_schedule
		&PositionRank{},              // position_rank
		&PositionGrade{},             // position_grade
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
