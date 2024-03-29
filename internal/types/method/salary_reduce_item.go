package method

import "gorm.io/gen"

type SalaryReduceItem interface {
	// select * from FN_C_SalaryReduceItemMultiply (@empIDs, @dateOnly1, @dateOnly2)
	QueryByEmployeeID(empIDs string, dateOnly1, dateOnly2 string) ([]gen.M, error)
}
