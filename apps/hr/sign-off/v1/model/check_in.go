package model

type CheckInApproveReq struct {
	UUID    string
	Comment string `json:"comment"`
}

type CheckInRejectReq struct {
	UUID    string
	Comment string `json:"comment"`
}
