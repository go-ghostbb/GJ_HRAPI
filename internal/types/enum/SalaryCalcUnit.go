package enum

import "database/sql/driver"

type SalaryCalcUnit string

const (
	DayCalc  SalaryCalcUnit = "day"
	HourCalc SalaryCalcUnit = "hour"
)

func (e *SalaryCalcUnit) Scan(value interface{}) error {
	*e = SalaryCalcUnit(value.(string))
	return nil
}

func (e SalaryCalcUnit) Value() (driver.Value, error) {
	return string(e), nil
}
