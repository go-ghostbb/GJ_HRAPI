package enum

import "database/sql/driver"

type SalaryCycle string

const (
	SalaryHour  SalaryCycle = "hour"
	SalaryMonth SalaryCycle = "month"
)

func (e *SalaryCycle) Scan(value interface{}) error {
	*e = SalaryCycle(value.(string))
	return nil
}

func (e SalaryCycle) Value() (driver.Value, error) {
	return string(e), nil
}
