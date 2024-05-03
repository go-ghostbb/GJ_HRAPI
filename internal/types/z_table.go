package types

func Basic() []interface{} {
	return []interface{}{
		&ConfigMap{},                 // config
		&CheckInRequestForm{},        // check_in_request_form
		&CheckInRequestFormDetail{},  // check_in_request_form_detail
		&CheckInSignOffFlow{},        // check_in_sign_off_flow
		&CheckInSignOffSetting{},     // check_in_sign_off_setting
		&CheckInStatus{},             // check_in_status
		&Employee{},                  // employee
		&LoginInformation{},          // login_information
		&Role{},                      // role
		&Permission{},                // permission
		&Menu{},                      // menu
		&Department{},                // department
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
		&LeaveRequestForm{},          // leave_request_form
		&LeaveSignOffFlow{},          // leave_sign_off_flow
		&LeaveCorrect{},              // leave_correct
		&OvertimeCheckIn{},           // overtime_check_in
		&OvertimeRequestForm{},       // overtime_request_form
		&OvertimeSignOffFlow{},       // overtime_sign_off_flow
		&OvertimeRequestFormRate{},   // overtime_request_form_rate
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
