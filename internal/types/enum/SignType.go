package enum

import "database/sql/driver"

type SignType string

const (
	DepartmentManager SignType = "department manager"
	SpecificEmployee  SignType = "specific employee"
	Agent             SignType = "agent"
)

func (l *SignType) Scan(value interface{}) error {
	*l = SignType(value.(string))
	return nil
}

func (l SignType) Value() (driver.Value, error) {
	return string(l), nil
}
