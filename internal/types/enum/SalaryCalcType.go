package enum

import "database/sql/driver"

type SalaryCalcType string

const (
	DefaultCalc   SalaryCalcType = "default"
	UnitCalc      SalaryCalcType = "unit"
	ConditionCalc SalaryCalcType = "condition"
)

func (e *SalaryCalcType) Scan(value interface{}) error {
	*e = SalaryCalcType(value.(string))
	return nil
}

func (e SalaryCalcType) Value() (driver.Value, error) {
	return string(e), nil
}
