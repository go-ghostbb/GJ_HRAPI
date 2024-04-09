package method

import "gorm.io/gen"

type Employee interface {
	// select id, card_number from @@table where card_number in (@cardNums);
	QueryIDWhereCardNum(cardNums []string) ([]gen.M, error)
	// select * from FN_C_EmployeeSeniority()
	CalculateEmployeeSeniority() ([]gen.M, error)
	// select * from FN_C_EmployeeSeniorityWithEndDate(@dateOnly)
	CalculateEmployeeSeniorityWithEndDate(dateOnly string) ([]gen.M, error)
}
