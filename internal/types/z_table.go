package types

func Basic() []interface{} {
	return []interface{}{
		&ConfigMap{},                 // config
		&CheckInStatus{},             // check_in_status
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
		&LeaveSignOffSetting{},       // leave_sign_off_setting
		&OvertimeSignOffSetting{},    // overtime_sign_off_setting
		&SalaryAddItem{},             // salary_add_item
		&SalaryReduceItem{},          // salary_reduce_item
		&CalcSalary{},                // calc_salary
		&CalcSalaryEmployee{},        // calc_salary_employee
		&CalcSalaryAdd{},             // calc_salary_add
		&CalcSalaryReduce{},          // calc_salary_reduce
	}
}

func M2M() []interface{} {
	return []interface{}{
		&RoleMenu{},
		&RolePermission{},
		&LeaveGroupEmployee{},
		&VacationGroupEmployee{},
		&SalaryAddItemEmployee{},
		&SalaryReduceItemEmployee{},
	}
}
