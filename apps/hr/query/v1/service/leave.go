package service

import (
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
)

func Leave(ctx context.Context) ILeave {
	return &leave{ctx: ctx}
}

type (
	ILeave interface {
		// GetLeaveCorrect 查詢可用假別
		GetLeaveCorrect(in model.GetLeaveCorrectReq) (out []*model.GetLeaveCorrectRes, err error)
	}

	leave struct {
		ctx context.Context
	}
)

// GetLeaveCorrect 查詢可用假別
func (l *leave) GetLeaveCorrect(in model.GetLeaveCorrectReq) (out []*model.GetLeaveCorrectRes, err error) {
	var (
		qCorrect = query.LeaveCorrect
		queryRes []*types.LeaveCorrect
	)

	queryRes, err = qCorrect.WithContext(l.ctx).
		Where(qCorrect.EmployeeID.Eq(in.EmployeeID)).
		Where(qCorrect.WithContext(l.ctx).Where(qCorrect.Effective.Like(driver.NewString(in.Year + "%"))).Or(qCorrect.Expired.Like(driver.NewString(in.Year + "%")))).
		Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}

	// copier
	if err = copier.Copy(&out, queryRes); err != nil {
		return nil, err
	}

	return out, nil
}
