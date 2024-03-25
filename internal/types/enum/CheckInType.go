package enum

import "database/sql/driver"

type CheckInType string

const (
	Work    CheckInType = "work"
	OffWork CheckInType = "off work"
)

func (e *CheckInType) Scan(value interface{}) error {
	*e = CheckInType(value.(string))
	return nil
}

func (e CheckInType) Value() (driver.Value, error) {
	return string(e), nil
}
