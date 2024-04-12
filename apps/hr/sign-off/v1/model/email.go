package model

import "hrapi/internal/types"

type EmailOption struct {
	Form    *types.LeaveRequestForm
	Subject string
	Msg     string
	Err     error
}
