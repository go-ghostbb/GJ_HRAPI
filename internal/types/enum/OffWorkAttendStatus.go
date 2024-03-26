package enum

import "database/sql/driver"

type OffWorkAttendStatus string

const (
	OffWorkNormal    OffWorkAttendStatus = "normal"
	OffWorkNotSwiped OffWorkAttendStatus = "not swiped"
	OffWorkEarly     OffWorkAttendStatus = "early"
	OffWorkOffDay    OffWorkAttendStatus = "off day"
)

func (e *OffWorkAttendStatus) Scan(value interface{}) error {
	*e = OffWorkAttendStatus(value.(string))
	return nil
}

func (e OffWorkAttendStatus) Value() (driver.Value, error) {
	return string(e), nil
}
