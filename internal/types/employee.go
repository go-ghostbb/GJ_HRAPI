package types

import (
	"gorm.io/gorm"
	"hrapi/internal/types/enum"
	"time"
)

type Employee struct {
	gorm.Model
	HireDate         time.Time             `gorm:"type:date;comment:到職日期" json:"hire_date"`
	TerminationDate  time.Time             `gorm:"type:date;comment:離職日期" json:"termination_date"`
	EmploymentStatus enum.EmploymentStatus `gorm:"size:20;default:active;comment:在職狀態，記錄員工的在職狀態，例如在職、停薪留職、已離職等" json:"employment_status"`
	Backend          bool                  `gorm:"default:false;not null;comment:是否可登入後台" json:"backend" v:"required"`

	// ---------基本訊息---------
	RealName   string    `gorm:"size:100;not null;comment:真實姓名" json:"real_name" v:"required"`
	NationalID string    `gorm:"size:50;not null;uniqueIndex;comment:身份證號碼" json:"national_id" v:"required"`
	Birth      time.Time `gorm:"type:date;comment:生日" json:"birth"`
	Email      string    `gorm:"size:100;comment:信箱" json:"email"`
	Mobile     string    `gorm:"size:20;comment:手機號" json:"mobile"`
	Avatar     string    `gorm:"comment:頭像" json:"avatar"`

	// -------帳號密碼資訊-------
	LoginInformation *LoginInformation `json:"login_information"`

	// -----角色(控制權限用)-----
	Roles []*Role `gorm:"many2many:employee_role;" json:"roles"`

	// --------部門--------
	DepartmentId uint        `gorm:"default:0;not null;comment:部門ID" json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentId;comment:使用者部門" json:"department"`
}

func (e *Employee) TableName() string {
	return "employee"
}
