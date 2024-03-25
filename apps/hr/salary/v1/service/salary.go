package service

import (
	"context"
	"hrapi/apps/hr/salary/v1/model"
)

func Salary(ctx context.Context) ISalary {
	return &salary{ctx: ctx}
}

type (
	ISalary interface {
	}

	salary struct {
		ctx context.Context
	}
)

func (s *salary) Calc(in model.SalaryCalcReq) (out model.SalaryCalcRes, err error) {
	return model.SalaryCalcRes{}, err
}
