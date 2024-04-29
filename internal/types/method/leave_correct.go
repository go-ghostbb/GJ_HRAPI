package method

import "gorm.io/gen"

type LeaveCorrect interface {
	// exec P_C_LeaveCorrectUpdateCount @id
	UpdateCount(id uint) (gen.RowsAffected, error)
	// select distinct leave_correct_id from FN_Q_LeaveCorrectByDateRange(@empID, @leaveID, @start, @end)
	FindByDateRange(empID, leaveID uint, start, end string) ([]uint, error)
}
