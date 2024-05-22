package enum

import "database/sql/driver"

type LeaveCycle string

const (
	Default       LeaveCycle = "year"
	Annual        LeaveCycle = "annual"
	Calendar      LeaveCycle = "calendar"
	CalendarTwice LeaveCycle = "calendar_twice"
)

func (l *LeaveCycle) Scan(value interface{}) error {
	*l = LeaveCycle(value.(string))
	return nil
}

func (l LeaveCycle) Value() (driver.Value, error) {
	return string(l), nil
}
