package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/paginator"
)

func Leave(ctx context.Context, page ...*paginator.Pagination) ILeave {
	if len(page) != 0 {
		return &leave{page: page[0], ctx: ctx}
	}
	return &leave{ctx: ctx}
}

type (
	ILeave interface {
		// GetLeaveCorrect 查詢可用假別
		GetLeaveCorrect(in model.GetLeaveCorrectReq) (out []*model.GetLeaveCorrectRes, err error)
		// GetLeaveCorrectByKeyword 根據關鍵字查詢可用假別
		GetLeaveCorrectByKeyword(in model.GetLeaveCorrectByKeywordReq) (out model.GetLeaveCorrectByKeywordRes, err error)
		// UpdateLeaveCorrect 更新leave correct
		UpdateLeaveCorrect(in model.PutLeaveCorrectReq) (err error)
	}

	leave struct {
		page *paginator.Pagination
		ctx  context.Context
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

// GetLeaveCorrectByKeyword 根據關鍵字查詢可用假別
func (l *leave) GetLeaveCorrectByKeyword(in model.GetLeaveCorrectByKeywordReq) (out model.GetLeaveCorrectByKeywordRes, err error) {
	var (
		qCorrect  = query.LeaveCorrect
		qEmployee = query.Employee
		queryRes  []*types.LeaveCorrect
	)

	queryRes, err = qCorrect.WithContext(l.ctx).
		Join(qEmployee, qEmployee.ID.EqCol(qCorrect.EmployeeID)).
		Where(qCorrect.WithContext(l.ctx).Where(qEmployee.RealName.Like("%" + in.Keyword + "%")).Or(qEmployee.Code.Like("%" + in.Keyword + "%"))).
		Where(qCorrect.WithContext(l.ctx).Where(qCorrect.Effective.Like(driver.NewString(in.Year + "%"))).Or(qCorrect.Expired.Like(driver.NewString(in.Year + "%")))).
		Offset(l.page.Offset).Limit(l.page.Limit).
		Preload(field.Associations).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, queryRes); err != nil {
		return
	}

	// count
	out.Total, err = qCorrect.WithContext(l.ctx).
		Join(qEmployee, qEmployee.ID.EqCol(qCorrect.EmployeeID)).
		Where(qCorrect.WithContext(l.ctx).Where(qEmployee.RealName.Like("%" + in.Keyword + "%")).Or(qEmployee.Code.Like("%" + in.Keyword + "%"))).
		Where(qCorrect.WithContext(l.ctx).Where(qCorrect.Effective.Like(driver.NewString(in.Year + "%"))).Or(qCorrect.Expired.Like(driver.NewString(in.Year + "%")))).
		Count()
	if err != nil {
		return
	}

	return out, nil
}

// UpdateLeaveCorrect 更新leave correct
func (l *leave) UpdateLeaveCorrect(in model.PutLeaveCorrectReq) (err error) {
	var (
		qLeaveCorrect = query.LeaveCorrect
		leaveCorrect  *types.LeaveCorrect
	)

	leaveCorrect, err = qLeaveCorrect.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeaveCorrect.ID.Eq(in.ID)).First()
	if err != nil {
		return err
	}

	leaveCorrect.Effective = in.Effective
	leaveCorrect.Expired = in.Expired
	leaveCorrect.Available = in.Available

	return qLeaveCorrect.WithContext(dbcache.WithCtx(l.ctx)).Save(leaveCorrect)
}
