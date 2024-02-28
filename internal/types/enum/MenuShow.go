package enum

import "database/sql/driver"

type MenuShow string

const (
	Software MenuShow = "software"
	Web      MenuShow = "web"
)

func (e *MenuShow) Scan(value interface{}) error {
	*e = MenuShow(value.(string))
	return nil
}

func (e MenuShow) Value() (driver.Value, error) {
	return string(e), nil
}
