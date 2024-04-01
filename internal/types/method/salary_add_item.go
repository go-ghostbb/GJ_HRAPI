package method

import "gorm.io/gen"

type SalaryAddItem interface {
	// select * from FN_C_SalaryAddItemMultiply (@empIDs, @dateOnly1, @dateOnly2)
	QueryByEmployeeIDMultiply(empIDs string, dateOnly1, dateOnly2 string) ([]gen.M, error)
	// select sa.*, sae.use_custom, sae.custom_amount from salary_add_item_employee sae
	//     join salary_add_item sa on (sa.id = sae.salary_add_item_id)
	//     where sae.employee_id = @empID and sa.name like @keyword
	QueryByEmployeeID(empID uint, keyword string) ([]gen.M, error)
}
