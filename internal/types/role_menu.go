package types

type RoleMenu struct {
	MenuID uint `json:"menuId"`
	RoleID uint `json:"roleId"`
}

func (r *RoleMenu) TableName() string {
	return "role_menu"
}
