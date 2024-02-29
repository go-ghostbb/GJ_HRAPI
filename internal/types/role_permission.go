package types

type RolePermission struct {
	RoleID       uint `json:"roleId"`
	PermissionID uint `json:"permissionId"`
}

func (r *RolePermission) TableName() string {
	return "role_permission"
}
