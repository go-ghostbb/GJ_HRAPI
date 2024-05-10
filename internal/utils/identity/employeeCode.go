package identity

import (
	"context"
	"fmt"
	gbtype "ghostbb.io/gb/container/gb_type"
	gberror "ghostbb.io/gb/errors/gb_error"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"hrapi/internal/query"
)

var (
	EmployeeCode = gbtype.NewUint(1)
)

var (
	employeeCodePrefix string
)

func employeeCodeInit() error {
	var (
		qConfig   = query.ConfigMap
		qEmployee = query.Employee
		nowCode   string
		err       error
	)
	// 尋找employee code prefix
	err = qConfig.WithContext(context.Background()).Select(qConfig.Value).Where(qConfig.Key.Eq("EmployeeCodePrefix")).Scan(&employeeCodePrefix)
	if err != nil {
		return err
	}
	if employeeCodePrefix == "" {
		return gberror.New("config: EmployeeCodePrefix not found")
	}

	// 尋找employee code
	nowCode, err = qEmployee.WithContext(context.Background()).GetMaxCode()
	if err != nil {
		return err
	}
	if nowCode == "" {
		return gberror.New("max employee code not found")
	}
	EmployeeCode.Set(gbconv.Uint(gbstr.Replace(nowCode, employeeCodePrefix, "")))

	return nil
}

// GetEmployeeCode 取得員工編號，併發可用
func GetEmployeeCode() string {
	return employeeCodePrefix + fmt.Sprintf("%04d", EmployeeCode.Add(1))
}
