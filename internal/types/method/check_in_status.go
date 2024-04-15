package method

import "gorm.io/gen"

type CheckInStatus interface {
	// select * from @@table
	//     where work_check_in_date between @dateOnly1 and @dateOnly2 and
	//           employee_id = @empID and
	//  		 {{if abnormal}}
	//           	(work_attend_proc_status = 'not processed' or off_work_attend_proc_status = 'not processed') and
	// 			 {{end}}
	//           deleted_at is null
	QueryByDateRangeAndEmpID(empID uint, dateOnly1, dateOnly2 string, abnormal bool) ([]*gen.T, error)
	// delete @@table where work_check_in_date = @dateOnly
	DeleteByDate(dateOnly string) (gen.RowsAffected, error)
	// delete @@table where work_check_in_date between @dateOnly1 and @dateOnly2
	DeleteByDateRange(dateOnly1, dateOnly2 string) (gen.RowsAffected, error)
	// exec P_C_CheckInStatusUpdateStatus @dateOnly
	UpdateStatus(dateOnly string) error
	// update @@table
	// {{if isWork}}
	//     set work_attend_time = @timeOnly, updated_at = getdate()
	// {{else}}
	//     set off_work_attend_time = @timeOnly, updated_at = getdate()
	// {{end}}
	// where
	//     work_shift_id = (select id from work_shift where code = @workShiftCode) and
	//     employee_id = (select id from employee where card_number = @cardNum) and
	//     {{if isWork}}
	//         work_check_in_date = @dateOnly
	//     {{else}}
	//         off_work_check_in_date = @dateOnly
	//     {{end}}
	UpdateTime(timeOnly, dateOnly, workShiftCode, cardNum string, isWork bool) (gen.RowsAffected, error)
	// select sum(w.total_hours) - sum(c.absence_hours) - sum(c.leave_hours) from check_in_status c
	//     join work_shift w on (c.work_shift_id = w.id)
	//     join employee e on (c.employee_id = e.id)
	// where
	//     c.work_shift_id != 0 and
	//     work_check_in_date between @dateOnly1 and @dateOnly2 and
	//     c.employee_id = @empID and
	//     e.salary_cycle = 'hour'
	QueryTotalAttendHours(empID uint, dateOnly1, dateOnly2 string) (float32, error)
	// exec P_C_CheckInHourUpdateByLeave @empID, @leaveRequestFromID
	UpdateHourByLeave(empID, leaveRequestFromID uint) error
	// exec P_C_CheckInHourSubByLeave @empID, @leaveRequestFromID
	SubHourByLeave(empID, leaveRequestFromID uint) error
}
