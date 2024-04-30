package service

import (
	"context"
	"database/sql"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	"gorm.io/gen/field"
	"hrapi/apps/hr/sign-off/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/email"
	"time"
)

func Overtime(ctx context.Context) IOvertime {
	return &overtime{ctx: ctx}
}

type (
	IOvertime interface {
		// GetByUUID 根據uuid獲取加班單資訊
		GetByUUID(uuid string) (form *types.OvertimeRequestForm, err error)
		// Approve 核准
		Approve(in model.OvertimeApproveReq) error
		// Reject 駁回
		Reject(in model.OvertimeRejectReq) error
		// NextFlow 啟動下一個流程
		NextFlow(uuid string)

		sendEmail(option *model.EmailOption[types.OvertimeRequestForm])
	}

	overtime struct {
		ctx context.Context
	}
)

// GetByUUID 根據uuid獲取加班單資訊
func (o *overtime) GetByUUID(uuid string) (form *types.OvertimeRequestForm, err error) {
	var (
		qForm = query.OvertimeRequestForm
		qFlow = query.OvertimeSignOffFlow

		// select overtime_request_form_id from overtime_sign_off_flow where uuid = ?
		subQueryFormID = qFlow.WithContext(o.ctx).Select(qFlow.OvertimeRequestFormID).Where(qFlow.UUID.Eq(uuid))
	)

	form, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).
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
func (o *overtime) Approve(in model.OvertimeApproveReq) error {
	err := query.Q.Transaction(func(tx *query.Query) error {
		var (
			qFlow = tx.OvertimeSignOffFlow
			qForm = tx.OvertimeRequestForm

			// select overtime_request_form_id from overtime_sign_off_flow where uuid = ?
			subQueryFormID = qFlow.WithContext(o.ctx).Select(qFlow.OvertimeRequestFormID).Where(qFlow.UUID.Eq(in.UUID))

			err error
		)

		// 更新此流程的簽核時間和簽核狀態
		_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Where(qFlow.UUID.Eq(in.UUID)).
			UpdateSimple(qFlow.SignDate.Value(time.Now()), qFlow.Status.Value(enum.SignStatusApprove), qFlow.Comment.Value(in.Comment))
		if err != nil {
			return err
		}

		// 如果表單為送簽中，改為簽核中
		_, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).
			Where(qForm.Columns(qForm.ID).Eq(subQueryFormID), qForm.SignStatus.Eq(enum.SignStatusSending)).
			UpdateSimple(qForm.SignStatus.Value(enum.SignStatusSigning))
		if err != nil {
			return err
		}

		// commit
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		return err
	}

	// 前往下一流程
	// 開啟一個新的service
	// 使用goroutine
	go Overtime(gbctx.New()).NextFlow(in.UUID)

	return nil
}

// Reject 駁回
func (o *overtime) Reject(in model.OvertimeRejectReq) error {
	// 查詢此uuid的請假單資訊
	var form, err = o.GetByUUID(in.UUID)
	if err != nil {
		return err
	}

	// 開啟事務
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qFlow = tx.OvertimeSignOffFlow
			qForm = tx.OvertimeRequestForm

			err error
		)

		// 更新此流程的簽核時間和簽核狀態
		_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Where(qFlow.UUID.Eq(in.UUID)).
			UpdateSimple(qFlow.SignDate.Value(time.Now()), qFlow.Status.Value(enum.SignStatusReject), qFlow.Comment.Value(in.Comment))
		if err != nil {
			return err
		}

		// 更新此請假單的狀態
		_, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).Where(qForm.ID.Eq(form.ID)).
			UpdateSimple(qForm.SignStatus.Value(enum.SignStatusReject))
		if err != nil {
			return err
		}

		// commit
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
	if err != nil {
		return err
	}

	// send email
	go Overtime(gbctx.New()).sendEmail(&model.EmailOption[types.OvertimeRequestForm]{
		Form: form,
		// 最後一關為起單人的完成後通知
		// 直接通知最後一關
		UUID:    form.SignOffFlow[len(form.SignOffFlow)-1].UUID,
		Subject: fmt.Sprintf("駁回【加班單】- %s", form.Order),
		Msg:     "加班單駁回",
	})

	return nil
}

// NextFlow 啟動下一個流程
func (o *overtime) NextFlow(uuid string) {
	var (
		form *types.OvertimeRequestForm
		err  error
	)

	// 獲得表單內容
	form, err = o.GetByUUID(uuid)
	if err != nil {
		g.Log().Error(o.ctx, "query form by uuid error:", err)
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
			o.handleSignOffPlusNotify(form, flow.UUID)
			// 如果是簽核通知，直接離開迴圈
			// 因為已經找到下一個流程
			return
		case enum.NotifyOnly:
			// 僅通知
			// 不return, 因為通知完需直接進下一個流程
			o.handleNotifyOnly(form, flow.UUID)
		}
	}

	// 如果跑出迴圈，代表跑完所有流程
	err = o.handleSignDone(form)
	if err != nil {
		g.Log().Errorf(o.ctx, "make overtime order:%s signing completed, but error in updating database:%s", form.Order, err.Error())
	}
}

// 處理簽核通知
func (o *overtime) handleSignOffPlusNotify(form *types.OvertimeRequestForm, uuid string) {
	var (
		qFlow = query.OvertimeSignOffFlow
		err   error
	)

	// 更改流程狀態為簽核中
	_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(qFlow.Status.Value(enum.SignStatusSigning))
	if err != nil {
		g.Log().Error(o.ctx, "update sign-off form fail:", err.Error())
		return
	}

	// send email
	o.sendEmail(&model.EmailOption[types.OvertimeRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("待簽核【加班單】- %s", form.Order),
		Msg:     "加班單簽核",
	})
}

// 處理僅通知
func (o *overtime) handleNotifyOnly(form *types.OvertimeRequestForm, uuid string) {
	var (
		qFlow = query.OvertimeSignOffFlow
		err   error
	)

	// send email
	emailOption := &model.EmailOption[types.OvertimeRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("通知【加班單】- %s", form.Order),
		Msg:     "加班單通知",
	}
	o.sendEmail(emailOption)

	updateCol := qFlow.Status.Value(enum.SignStatusOnlyNotifySuc)

	if emailOption.Err != nil {
		// 如果有錯誤
		g.Log().Error(o.ctx, "send email fail:", emailOption.Err.Error())
		// 更新欄位為通知失敗
		updateCol = qFlow.Status.Value(enum.SignStatusOnlyNotifyFail)
	}

	// 更新流程狀態
	_, err = qFlow.WithContext(dbcache.WithCtx(o.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(updateCol)
	if err != nil {
		g.Log().Error(o.ctx, "update sign-off form fail:", err.Error())
	}
}

// 處理跑完所有流程後操作
func (o *overtime) handleSignDone(form *types.OvertimeRequestForm) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm = tx.OvertimeRequestForm
			err   error
		)

		// 更新表單狀態
		_, err = qForm.WithContext(dbcache.WithCtx(o.ctx)).Where(qForm.ID.Eq(form.ID)).UpdateSimple(qForm.SignStatus.Value(enum.SignStatusApprove))
		if err != nil {
			return err
		}

		// TODO: 更新狀態表

		// commit
		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelReadUncommitted,
	})
}

// 寄信通知
func (o *overtime) sendEmail(option *model.EmailOption[types.OvertimeRequestForm]) {
	var (
		qEmployee     = query.Employee
		employeeEmail string
		url           string
		body          string
		err           error
	)

	// 查詢員工資訊
	err = qEmployee.WithContext(o.ctx).Select(qEmployee.Email).Where(qEmployee.ID.Eq(option.Form.EmployeeID)).Scan(&employeeEmail)
	if err != nil {
		option.Err = gberror.Newf("query employee email error:%s", err.Error())
		return
	}

	// get url
	url, err = utils.GetURL("/sign-off/overtime")
	if err != nil {
		option.Err = gberror.Newf("get url error:%s", err.Error())
		return
	}

	url = fmt.Sprintf("%s?uuid=%s&locale=%s", url, option.UUID, option.Form.SignOffFlow[0].Locale)

	// template
	view := g.View("overtimeSign")
	body, err = view.ParseDefault(o.ctx, g.Map{
		"url":     url,
		"content": option.Form.Order,
		"msg":     option.Msg,
	})
	if err != nil {
		option.Err = gberror.Newf("parse template(mail/overtimeSign.html) error:%s", err.Error())
		return
	}

	// send
	err = email.Service.SendTo([]string{employeeEmail}, option.Subject, body)
	if err != nil {
		option.Err = gberror.Newf("send email error:%s", err.Error())
	}
}
