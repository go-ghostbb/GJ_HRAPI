package model

type OvertimeApproveReq struct {
	UUID    string
	Comment string `json:"comment"`
}

type OvertimeRejectReq struct {
	UUID    string
	Comment string `json:"comment"`
}
