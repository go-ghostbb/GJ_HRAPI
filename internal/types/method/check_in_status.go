package method

import "gorm.io/gen"

type CheckInStatus interface {
	// select * from FN_Q_CheckInStatusByDateRange(@empID, @dateOnly1, @dateOnly2, @abnormal);
	QueryByDateRangeAndEmpID(empID uint, dateOnly1, dateOnly2 string, abnormal bool) ([]*gen.T, error)
	// exec P_C_CheckInStatusUpdateStatus @startDate, @endDate, @empIDs
	UpdateStatus(startDate, endDate, empIDs string) error
	// exec P_C_CheckInStatusUpdateTime @datetime, @workShiftCode, @cardNum, @isWork
	UpdateTime(datetime, workShiftCode, cardNum string, isWork bool) error
	// select sum(w.total_hours) - sum(c.absence_hours) - sum(c.leave_hours) from check_in_status c
	//     join work_shift w on (c.work_shift_id = w.id)
	//     join employee e on (c.employee_id = e.id)
	// where
	//     c.work_shift_id != 0 and
	//     work_check_in_date between @dateOnly1 and @dateOnly2 and
	//     c.employee_id = @empID and
	//     e.salary_cycle = 'hour'
	QueryTotalAttendHours(empID uint, dateOnly1, dateOnly2 string) (float32, error)
}
