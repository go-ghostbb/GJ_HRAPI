package method

import "gorm.io/gen"

type SalaryAddItem interface {
	// select * from FN_C_SalaryAddItemMultiply (@empIDs, @dateOnly1, @dateOnly2)
	QueryByEmployeeID(empIDs string, dateOnly1, dateOnly2 string) ([]gen.M, error)
}
