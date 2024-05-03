package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type OvertimeRequestForm struct {
	gorm.Model
	Order          string          `gorm:"size:50;uniqueIndex;comment:單號" json:"order"`
	Date           driver.Date     `gorm:"type:date;not null;comment:日期" json:"date"`
	StartTime      driver.Time     `gorm:"type:time(0);not null;comment:開始時間" json:"startTime"`
	EndTime        driver.Time     `gorm:"type:time(0);not null;comment:結束時間" json:"endTime"`
	EstimatedHours float32         `gorm:"comment:預計時數" json:"estimatedHours"`
	SignStatus     enum.SignStatus `gorm:"type:tinyint;not null;default:0;comment:0 送簽中, 1 簽核中, 2 通過, 3 拒絕" json:"signStatus"`
	Remark         string          `gorm:"comment:備註" json:"remark"`

	// 休假日
	VacationID uint      `gorm:"comment:假別ID" json:"vacationId"`
	Vacation   *Vacation `gorm:"foreignKey:VacationID" json:"vacation"`

	// 加班人員
	EmployeeID uint      `gorm:"default:0;not null;comment:加班員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:加班員工" json:"employee"`

	// 加班人員部門
	DepartmentID uint        `gorm:"default:0;not null;comment:加班部門ID" json:"departmentId"`
	Department   *Department `gorm:"foreignKey:DepartmentID;comment:加班部門" json:"department"`

	// 簽核流程
	SignOffFlow []*OvertimeSignOffFlow `gorm:"foreignKey:OvertimeRequestFormID" json:"signOffFlow"`

	// 加班倍率
	Rate []*OvertimeRequestFormRate `gorm:"foreignKey:OvertimeRequestFormID" json:"rate"`
}

func (o *OvertimeRequestForm) TableName() string {
	return "overtime_request_form"
}
