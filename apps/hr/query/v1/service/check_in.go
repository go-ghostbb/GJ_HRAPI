package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
)

func CheckIn(ctx context.Context, page ...*paginator.Pagination) ICheckIn {
	if len(page) > 0 {
		return &checkIn{ctx, page[0]}
	}
	return &checkIn{ctx: ctx}
}

type (
	ICheckIn interface {
		// QueryStatus 查詢出勤狀態
		QueryStatus(in model.QueryStatusReq) (out []*model.QueryStatusRes, err error)
	}

	checkIn struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// QueryStatus 查詢出勤狀態
func (c *checkIn) QueryStatus(in model.QueryStatusReq) (out []*model.QueryStatusRes, err error) {
	var (
		qCheckInStatus = query.CheckInStatus
		result         []*types.CheckInStatus
	)

	result, err = qCheckInStatus.WithContext(c.ctx).
		Preload(field.Associations, qCheckInStatus.Employee.Department).
		QueryByDateRangeAndEmpID(in.EmployeeID, in.DateRangeStart, in.DateRangeEnd, in.Abnormal)
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&out, result); err != nil {
		return nil, err
	}

	return out, nil
}
