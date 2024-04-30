package model

import "hrapi/internal/types"

type EmailOption[T types.LeaveRequestForm | types.CheckInRequestForm | types.OvertimeRequestForm] struct {
	Form    *T
	UUID    string
	Subject string
	Msg     string
	Err     error
}
