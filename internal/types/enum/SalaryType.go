package enum

import "database/sql/driver"

type SalaryType string

const (
	FixedSalary    SalaryType = "fixed"
	NonFixedSalary SalaryType = "non-fixed"
)

func (e *SalaryType) Scan(value interface{}) error {
	*e = SalaryType(value.(string))
	return nil
}

func (e SalaryType) Value() (driver.Value, error) {
	return string(e), nil
}
