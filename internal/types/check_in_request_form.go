package types

import (
	gbstr "ghostbb.io/gb/text/gb_str"
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
)

type CheckInRequestForm struct {
	gorm.Model
	Order      string          `gorm:"size:50;uniqueIndex;comment:單號" json:"order"`
	SignStatus enum.SignStatus `gorm:"type:tinyint;not null;default:0;comment:0 送簽中, 1 簽核中, 2 通過, 3 拒絕" json:"signStatus"`
	Attach     string          `gorm:"comment:附件" json:"attach"`
	Remark     string          `gorm:"comment:備註" json:"remark"`

	// 補打卡人員
	EmployeeID uint      `gorm:"default:0;not null;comment:請假員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:請假員工" json:"employee"`

	// 補打卡人員部門
	DepartmentID uint        `gorm:"default:0;not null;comment:請假部門ID" json:"departmentId"`
	Department   *Department `gorm:"foreignKey:DepartmentID;comment:請假部門" json:"department"`

	// 簽核流程
	SignOffFlow []*CheckInSignOffFlow `gorm:"foreignKey:CheckInRequestFormID" json:"signOffFlow"`

	// 明細
	Detail []*CheckInRequestFormDetail `gorm:"foreignKey:CheckInRequestFormID" json:"detail"`
}

func (c *CheckInRequestForm) TableName() string {
	return "check_in_request_form"
}

// AttachArray for copier
func (c *CheckInRequestForm) AttachArray() []string {
	result := gbstr.Split(c.Attach, ",")
	// Remove empty string
	for index, value := range result {
		if value == "" {
			result = append(result[:index], result[index+1:]...)
		}
	}
	return result
}
