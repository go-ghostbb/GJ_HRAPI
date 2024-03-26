package enum

import "database/sql/driver"

type WorkAttendStatus string

const (
	WorkNormal    WorkAttendStatus = "normal"
	WorkNotSwiped WorkAttendStatus = "not swiped"
	WorkLate      WorkAttendStatus = "late"
	WorkOffDay    WorkAttendStatus = "off day"
)

func (e *WorkAttendStatus) Scan(value interface{}) error {
	*e = WorkAttendStatus(value.(string))
	return nil
}

func (e WorkAttendStatus) Value() (driver.Value, error) {
	return string(e), nil
}
