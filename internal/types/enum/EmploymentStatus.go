package enum

import "database/sql/driver"

type EmploymentStatus string

const (
	Active      EmploymentStatus = "active"
	UnpaidLeave EmploymentStatus = "unpaid leave"
	Resigned    EmploymentStatus = "resigned"
)

func (e *EmploymentStatus) Scan(value interface{}) error {
	*e = EmploymentStatus(value.(string))
	return nil
}

func (e EmploymentStatus) Value() (driver.Value, error) {
	return string(e), nil
}
