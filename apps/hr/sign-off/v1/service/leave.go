package service

import (
	"context"
	"database/sql"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"gorm.io/gen/field"
	"hrapi/apps/hr/sign-off/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/email"
	"time"
)

func Leave(ctx context.Context) ILeave {
	return &leave{ctx: ctx}
}

type (
	ILeave interface {
		// GetByUUID 根據uuid查詢請假單
		GetByUUID(uuid string) (form *types.LeaveRequestForm, err error)
		// Approve 核准
		Approve(in model.LeaveApproveReq) error
		// Reject 駁回
		Reject(in model.LeaveRejectReq) error

		// 判斷下一流程
		nextFlow(uuid string)
		// 寄信
		sendEmail(option *model.EmailOption[types.LeaveRequestForm])
	}

	leave struct {
		ctx context.Context
	}
)

// GetByUUID 根據uuid查詢請假單
func (l *leave) GetByUUID(uuid string) (form *types.LeaveRequestForm, err error) {
	var (
		qForm = query.LeaveRequestForm
		qFlow = query.LeaveSignOffFlow

		// select leave_request_form_id from leave_sign_off_flow where uuid = ?
		subQueryFormID = qFlow.WithContext(l.ctx).Select(qFlow.LeaveRequestFormID).Where(qFlow.UUID.Eq(uuid))
	)

	form, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).
		Where(qForm.Columns(qForm.ID).Eq(subQueryFormID)).
		Preload(
			field.Associations,
			qForm.SignOffFlow.SignOffEmployee,
		).First()
	if err != nil {
		return nil, err
	}

	return
}

// Approve 核准
func (l *leave) Approve(in model.LeaveApproveReq) error {
	// 開啟事務
	err := query.Q.Transaction(func(tx *query.Query) error {
		var (
			qFlow = tx.LeaveSignOffFlow
			qForm = tx.LeaveRequestForm

			// select leave_request_form_id from leave_sign_off_flow where uuid = ?
			subQueryFormID = qFlow.WithContext(l.ctx).Select(qFlow.LeaveRequestFormID).Where(qFlow.UUID.Eq(in.UUID))

			err error
		)

		// 更新此流程的簽核時間和簽核狀態
		_, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.UUID.Eq(in.UUID)).
			UpdateSimple(qFlow.SignDate.Value(time.Now()), qFlow.Status.Value(enum.SignStatusApprove), qFlow.Comment.Value(in.Comment))
		if err != nil {
			return err
		}

		// 如果表單為送簽中，改為簽核中
		_, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).
			Where(qForm.Columns(qForm.ID).Eq(subQueryFormID), qForm.SignStatus.Eq(enum.SignStatusSending)).
			UpdateSimple(qForm.SignStatus.Value(enum.SignStatusSigning))
		if err != nil {
			return err
		}

		// commit
		return nil
	})
	if err != nil {
		return err
	}

	// 前往下一流程
	// 開啟一個新的service
	// 使用goroutine
	go Leave(gbctx.New()).nextFlow(in.UUID)

	return nil
}

// Reject 駁回
func (l *leave) Reject(in model.LeaveRejectReq) error {
	// 查詢此uuid的請假單資訊
	var form, err = l.GetByUUID(in.UUID)
	if err != nil {
		return err
	}

	// 開啟事務
	opts := &sql.TxOptions{
		// 啟用讀取未提交
		Isolation: sql.LevelReadCommitted,
	}
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qFlow          = tx.LeaveSignOffFlow
			qForm          = tx.LeaveRequestForm
			qCheckInStatus = tx.CheckInStatus

			err error
		)

		// 更新此流程的簽核時間和簽核狀態
		_, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.UUID.Eq(in.UUID)).
			UpdateSimple(qFlow.SignDate.Value(time.Now()), qFlow.Status.Value(enum.SignStatusReject), qFlow.Comment.Value(in.Comment))
		if err != nil {
			return err
		}

		// 更新此請假單的狀態
		_, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).Where(qForm.ID.Eq(form.ID)).
			UpdateSimple(qForm.SignStatus.Value(enum.SignStatusReject))
		if err != nil {
			return err
		}

		// 更新遞延表或可用請假表
		err = l.updateLeaveCorrect(tx, form)
		if err != nil {
			return err
		}

		// 更新出勤狀態表
		// 查詢form
		err = qCheckInStatus.WithContext(dbcache.WithCtx(l.ctx)).UpdateStatus(form.StartDate.Format(), form.EndDate.Format(), gbconv.String(form.EmployeeID))
		if err != nil {
			return err
		}

		// commit
		return nil
	}, opts)
	if err != nil {
		return err
	}

	// send email
	go Leave(gbctx.New()).sendEmail(&model.EmailOption[types.LeaveRequestForm]{
		Form: form,
		// 最後一關為起單人的完成後通知
		// 直接通知最後一關
		UUID:    form.SignOffFlow[len(form.SignOffFlow)-1].UUID,
		Subject: fmt.Sprintf("駁回【請假單】- %s", form.Order),
		Msg:     "請假單駁回",
	})

	return nil
}

// nextFlow 進行下一個流程
func (l *leave) nextFlow(uuid string) {
	var (
		form *types.LeaveRequestForm
		err  error
	)

	// 獲得表單內容
	form, err = l.GetByUUID(uuid)
	if err != nil {
		g.Log().Error(l.ctx, "query form by uuid error:", err)
		return
	}

	// 開始遍歷流程
	for _, flow := range form.SignOffFlow {
		// 如果status == SignStatusSending(送簽中)代表流程未開啟
		// 搜尋到代表為下一個流程
		if flow.Status != enum.SignStatusSending {
			continue
		}

		switch flow.Notify {
		case enum.SignOffPlusNotify:
			// 簽核通知
			l.handleSignOffPlusNotify(form, flow.UUID)
			// 如果是簽核通知，直接離開迴圈
			// 因為已經找到下一個流程
			return
		case enum.NotifyOnly:
			// 僅通知
			// 不return, 因為通知完需直接進下一個流程
			l.handleNotifyOnly(form, flow.UUID)
		}
	}

	// 如果跑出迴圈，代表跑完所有流程
	err = l.handleSignDone(form)
	if err != nil {
		g.Log().Errorf(l.ctx, "leave order:%s signing completed, but error in updating database:%s", form.Order, err.Error())
	}
}

// 處理簽核通知
func (l *leave) handleSignOffPlusNotify(form *types.LeaveRequestForm, uuid string) {
	var (
		qFlow = query.LeaveSignOffFlow
		err   error
	)

	// 更改流程狀態為簽核中
	_, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(qFlow.Status.Value(enum.SignStatusSigning), qFlow.SignDate.Value(time.Now()))
	if err != nil {
		g.Log().Error(l.ctx, "update sign-off form fail:", err.Error())
		return
	}

	// send email
	l.sendEmail(&model.EmailOption[types.LeaveRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("待簽核【請假單】- %s", form.Order),
		Msg:     "請假單簽核",
	})
}

// 處理僅通知
func (l *leave) handleNotifyOnly(form *types.LeaveRequestForm, uuid string) {
	var (
		qFlow = query.LeaveSignOffFlow
		err   error
	)

	// send email
	emailOption := &model.EmailOption[types.LeaveRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("通知【請假單】- %s", form.Order),
		Msg:     "請假單通知",
	}
	l.sendEmail(emailOption)

	updateCol := qFlow.Status.Value(enum.SignStatusOnlyNotifySuc)

	if emailOption.Err != nil {
		// 如果有錯誤
		g.Log().Error(l.ctx, "send email fail:", emailOption.Err.Error())
		// 更新欄位為通知失敗
		updateCol = qFlow.Status.Value(enum.SignStatusOnlyNotifyFail)
	}

	// 更新流程狀態
	_, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(updateCol)
	if err != nil {
		g.Log().Error(l.ctx, "update sign-off form fail:", err.Error())
	}
}

// 處理跑完所有流程後操作
func (l *leave) handleSignDone(form *types.LeaveRequestForm) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm          = tx.LeaveRequestForm
			qCheckInStatus = tx.CheckInStatus
			err            error
		)

		// 更新表單狀態
		_, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).Where(qForm.ID.Eq(form.ID)).UpdateSimple(qForm.SignStatus.Value(enum.SignStatusApprove))
		if err != nil {
			return err
		}

		// 如果不是從遞延表裡面使用的話
		err = l.updateLeaveCorrect(tx, form)
		if err != nil {
			return err
		}

		// 更新出勤狀態表
		err = qCheckInStatus.WithContext(dbcache.WithCtx(l.ctx)).
			UpdateStatus(
				form.StartDate.AddDate(0, 0, -1).Format(),
				form.EndDate.AddDate(0, 0, 1).Format(),
				gbconv.String(form.EmployeeID),
			)
		if err != nil {
			return err
		}

		// commit
		return nil
	}, &sql.TxOptions{
		// 啟用讀取未提交
		Isolation: sql.LevelReadCommitted,
	})
}

// 更新請假可用表
func (l *leave) updateLeaveCorrect(tx *query.Query, form *types.LeaveRequestForm) error {
	var (
		qLeaveCorrect = tx.LeaveCorrect
		lcIDs         = make([]uint, 0)
		err           error
	)

	// 查詢表單內日期區間對應的所有可用表id
	lcIDs, err = qLeaveCorrect.WithContext(l.ctx).FindByDateRange(form.EmployeeID, form.LeaveID, form.StartDate.Format(), form.EndDate.Format())
	if err != nil {
		return err
	}

	// 如果有一個id為0
	// 代表請假有到下一個週期，且人事未結算本週期
	for _, id := range lcIDs {
		if id == 0 {
			return gberror.New("leave_correct id is 0")
		}
	}

	// 更新可用表
	var rowCount int64
	for _, id := range lcIDs {
		rowCount, err = qLeaveCorrect.WithContext(l.ctx).UpdateCount(id)
		if err != nil {
			return err
		}
		if rowCount == 0 {
			return gberror.New("leave_correct update quantity is 0")
		}
	}

	return nil
}

// 寄信通知
func (l *leave) sendEmail(option *model.EmailOption[types.LeaveRequestForm]) {
	var (
		qEmployee     = query.Employee
		employeeEmail string
		url           string
		body          string
		err           error
	)

	// 查詢員工資訊
	err = qEmployee.WithContext(l.ctx).Select(qEmployee.Email).Where(qEmployee.ID.Eq(option.Form.EmployeeID)).Scan(&employeeEmail)
	if err != nil {
		option.Err = gberror.Newf("query employee email error:%s", err.Error())
		return
	}

	// get url
	url, err = utils.GetURL("/sign-off/leave")
	if err != nil {
		option.Err = gberror.Newf("get url error:%s", err.Error())
		return
	}

	url = fmt.Sprintf("%s?uuid=%s&locale=%s", url, option.UUID, option.Form.SignOffFlow[0].Locale)

	// template
	view := g.View("leaveSign")
	body, err = view.ParseDefault(l.ctx, g.Map{
		"url":     url,
		"content": option.Form.Order,
		"msg":     option.Msg,
	})
	if err != nil {
		option.Err = gberror.Newf("parse template(mail/leaveSign.html) error:%s", err.Error())
		return
	}

	// send
	err = email.Service.SendTo([]string{employeeEmail}, option.Subject, body)
	if err != nil {
		option.Err = gberror.Newf("send email error:%s", err.Error())
	}
}
