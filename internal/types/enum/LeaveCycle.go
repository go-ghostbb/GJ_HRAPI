package enum

import "database/sql/driver"

type LeaveCycle string

const (
	Default  LeaveCycle = "year"
	Annual   LeaveCycle = "half"
	Calendar LeaveCycle = "calendar"
)

func (l *LeaveCycle) Scan(value interface{}) error {
	*l = LeaveCycle(value.(string))
	return nil
}

func (l LeaveCycle) Value() (driver.Value, error) {
	return string(l), nil
}
