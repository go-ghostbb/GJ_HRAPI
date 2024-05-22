package method

import "gorm.io/gen"

type LeaveCorrect interface {
	// {{if signing}}
	//   exec P_C_LeaveCorrectUpdateByFormSigning @leaveRequestFormID
	// {{else}}
	// 	 exec P_C_LeaveCorrectUpdateByFormApproveOrReject @leaveRequestFormID
	// {{end}}
	UpdateCount(leaveRequestFormID uint, signing bool) (gen.RowsAffected, error)
	// select distinct leave_correct_id from FN_Q_LeaveCorrectByDateRange(@empID, @leaveID, @start, @end)
	FindByDateRange(empID, leaveID uint, start, end string) ([]uint, error)
}
