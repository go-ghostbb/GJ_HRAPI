package method

import "gorm.io/gen"

type Employee interface {
	// select id, card_number from @@table where card_number in (@cardNums);
	QueryIDWhereCardNum(cardNums []string) ([]gen.M, error)
	// select * from FN_C_EmployeeSeniority()
	CalculateEmployeeSeniority() ([]gen.M, error)
	// select * from FN_C_EmployeeSeniorityWithEndDate(@dateOnly)
	CalculateEmployeeSeniorityWithEndDate(dateOnly string) ([]gen.M, error)
	// select max(code)
	// from employee
	// where
	//     exists (select value from config_map where [key] = 'EmployeeCodePrefix' and deleted_at is null) and
	//     code like (select concat(value, '%') from config_map where [key] = 'EmployeeCodePrefix' and deleted_at is null)
	GetMaxCode() (string, error)
}
