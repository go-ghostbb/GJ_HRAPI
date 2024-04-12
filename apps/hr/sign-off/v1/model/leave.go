package model

type LeaveApproveReq struct {
	UUID    string
	Comment string `json:"comment"`
}

type LeaveRejectReq struct {
	UUID    string
	Comment string `json:"comment"`
}
