package model

type FilingCheckInStatusReq struct {
	EmployeeID []uint   `json:"employeeId"`
	DateRange  []string `json:"dateRange" binding:"required"`
}

type CheckInData struct {
	CardNumber  string `json:"cardNumber" binding:"required"`
	DateTime    string `json:"dateTime" binding:"required"`
	CheckInType string `json:"checkInType" binding:"required"`
}

type UploadDataReq struct {
	*CheckInData
}
