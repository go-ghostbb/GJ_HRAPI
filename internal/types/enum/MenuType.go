package enum

import "database/sql/driver"

type MenuType string

const (
	Directory MenuType = "directory"
	Menu      MenuType = "menu"
	IFrame    MenuType = "iframe"
)

func (e *MenuType) Scan(value interface{}) error {
	*e = MenuType(value.(string))
	return nil
}

func (e MenuType) Value() (driver.Value, error) {
	return string(e), nil
}
