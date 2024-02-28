package types

import "gorm.io/gorm"

type LoginInformation struct {
	gorm.Model
	EmployeeID uint   `gorm:"unique;index;not null" json:"employee_id"`
	Status     bool   `gorm:"not null;default:true;comment:狀態" json:"status"`
	Account    string `gorm:"size:50;not null;unique;comment:帳號" json:"account"`
	Password   string `gorm:"size:150;not null;comment:密碼" json:"password"`

	Employee *Employee `gorm:"foreignKey:EmployeeID;comment:員工資訊" json:"employee"`
}

func (l *LoginInformation) TableName() string {
	return "login_information"
}
