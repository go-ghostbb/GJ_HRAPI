package types

import (
	gbstr "ghostbb.io/gb/text/gb_str"
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
)

type LeaveRequestForm struct {
	gorm.Model
	Order            string          `gorm:"size:50;uniqueIndex;comment:單號" json:"order"`
	StartDate        driver.Date     `gorm:"type:date;not null;comment:起始日期" json:"startDate"`
	StartTime        driver.Time     `gorm:"type:time(0);not null;comment:起始時間" json:"startTime"`
	EndDate          driver.Date     `gorm:"type:date;not null;comment:結束日期" json:"endDate"`
	EndTime          driver.Time     `gorm:"type:time(0);not null;comment:結束時間" json:"endTime"`
	Remark           string          `gorm:"comment:備註" json:"remark"`
	SignStatus       enum.SignStatus `gorm:"type:tinyint;not null;default:0;comment:0 送簽中, 1 簽核中, 2 通過, 3 拒絕" json:"signStatus"`
	LeaveMinuteCount float32         `gorm:"comment:請假分鐘數" json:"leaveMinuteCount"`
	LeaveHourCount   float32         `gorm:"comment:請假時數" json:"leaveHourCount"`
	LeaveDayCount    float32         `gorm:"comment:請假天數" json:"leaveDayCount"`
	IsDefer          bool            `gorm:"comment:是否使用遞延表" json:"isDefer"`
	Attach           string          `gorm:"comment:附件" json:"attach"`
	// 假別
	LeaveID uint   `gorm:"default:0;not null;comment:假別ID" json:"leaveId"`
	Leave   *Leave `gorm:"foreignKey:LeaveID;comment:假別" json:"leave"`

	// 請假人員
	EmployeeID uint      `gorm:"default:0;not null;comment:請假員工ID" json:"employeeId"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;comment:請假員工" json:"employee"`

	// 請假部門
	DepartmentID uint        `gorm:"default:0;not null;comment:請假部門ID" json:"departmentId"`
	Department   *Department `gorm:"foreignKey:DepartmentID;comment:請假部門" json:"department"`

	// 代理人員
	ProxyEmployeeID uint      `gorm:"default:0;not null;comment:代理員工ID" json:"proxyEmployeeId"`
	ProxyEmployee   *Employee `gorm:"foreignKey:ProxyEmployeeID;comment:代理員工" json:"proxyEmployee"`

	// 代理人員部門
	ProxyDepartmentID uint        `gorm:"default:0;not null;comment:代理員工部門ID" json:"proxyDepartmentId"`
	ProxyDepartment   *Department `gorm:"foreignKey:ProxyDepartmentID;comment:代理員工部門" json:"proxyDepartment"`

	// 簽核流程
	SignOffFlow []*LeaveSignOffFlow `gorm:"foreignKey:LeaveRequestFormID" json:"signOffFlow"`
}

func (l *LeaveRequestForm) TableName() string {
	return "leave_request_form"
}

// DepartmentName for copier
func (l *LeaveRequestForm) DepartmentName() string {
	return l.Department.Name
}

// ProxyEmployeeRealName for copier
func (l *LeaveRequestForm) ProxyEmployeeRealName() string {
	return l.ProxyEmployee.RealName
}

// ProxyDepartmentName for copier
func (l *LeaveRequestForm) ProxyDepartmentName() string {
	return l.ProxyDepartment.Name
}

// LeaveName for copier
func (l *LeaveRequestForm) LeaveName() string {
	return l.Leave.Name
}

// LeaveMinimum for copier
func (l *LeaveRequestForm) LeaveMinimum() uint {
	return l.Leave.Minimum
}

// DateArray for copier
func (l *LeaveRequestForm) DateArray() []string {
	return []string{l.StartDate.Format(), l.EndDate.Format()}
}

// AttachArray for copier
func (l *LeaveRequestForm) AttachArray() []string {
	result := gbstr.Split(l.Attach, ",")
	// Remove empty string
	for index, value := range result {
		if value == "" {
			result = append(result[:index], result[index+1:]...)
		}
	}
	return result
}
