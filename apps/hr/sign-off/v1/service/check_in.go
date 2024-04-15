package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	"gorm.io/gen/field"
	"hrapi/apps/hr/sign-off/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/email"
)

func CheckIn(ctx context.Context) ICheckIn {
	return &checkIn{ctx: ctx}
}

type (
	ICheckIn interface {
		// GetByUUID 根據uuid查詢補打卡單
		GetByUUID(uuid string) (form *types.CheckInRequestForm, err error)

		// NextFlow 下一個流程
		NextFlow(uuid string)
	}

	checkIn struct {
		ctx context.Context
	}
)

// GetByUUID 根據uuid查詢補打卡單
func (c *checkIn) GetByUUID(uuid string) (form *types.CheckInRequestForm, err error) {
	var (
		qForm = query.CheckInRequestForm
		qFlow = query.CheckInSignOffFlow

		// select leave_request_form_id from leave_sign_off_flow where uuid = ?
		subQueryFormID = qFlow.WithContext(c.ctx).Select(qFlow.CheckInRequestFormID).Where(qFlow.UUID.Eq(uuid))
	)

	form, err = qForm.WithContext(dbcache.WithCtx(c.ctx)).
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

// NextFlow 下一個流程
func (c *checkIn) NextFlow(uuid string) {
	var (
		form *types.CheckInRequestForm
		err  error
	)

	// 獲得表單內容
	form, err = c.GetByUUID(uuid)
	if err != nil {
		g.Log().Error(c.ctx, "query form by uuid error:", err)
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
			c.handleSignOffPlusNotify(form, flow.UUID)
			// 如果是簽核通知，直接離開迴圈
			// 因為已經找到下一個流程
			return
		case enum.NotifyOnly:
			// 僅通知
			// 不return, 因為通知完需直接進下一個流程
			c.handleNotifyOnly(form, flow.UUID)
		}
	}

	// 如果跑出迴圈，代表跑完所有流程
	err = c.handleSignDone(form)
	if err != nil {
		g.Log().Errorf(c.ctx, "make check in order:%s signing completed, but error in updating database:%s", form.Order, err.Error())
	}
}

// 處理簽核通知
func (c *checkIn) handleSignOffPlusNotify(form *types.CheckInRequestForm, uuid string) {
	var (
		qFlow = query.CheckInSignOffFlow
		err   error
	)

	// 更改流程狀態為簽核中
	_, err = qFlow.WithContext(dbcache.WithCtx(c.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(qFlow.Status.Value(enum.SignStatusSigning))
	if err != nil {
		g.Log().Error(c.ctx, "update sign-off form fail:", err.Error())
		return
	}

	// send email
	c.sendEmail(&model.EmailOption[types.CheckInRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("待簽核【請假單】- %s", form.Order),
		Msg:     "請假單簽核",
	})
}

// 處理僅通知
func (c *checkIn) handleNotifyOnly(form *types.CheckInRequestForm, uuid string) {
	var (
		qFlow = query.CheckInSignOffFlow
		err   error
	)

	// send email
	emailOption := &model.EmailOption[types.CheckInRequestForm]{
		Form:    form,
		UUID:    uuid,
		Subject: fmt.Sprintf("通知【請假單】- %s", form.Order),
		Msg:     "請假單通知",
	}
	c.sendEmail(emailOption)

	updateCol := qFlow.Status.Value(enum.SignStatusOnlyNotifySuc)

	if emailOption.Err != nil {
		// 如果有錯誤
		g.Log().Error(c.ctx, "send email fail:", emailOption.Err.Error())
		// 更新欄位為通知失敗
		updateCol = qFlow.Status.Value(enum.SignStatusOnlyNotifyFail)
	}

	// 更新流程狀態
	_, err = qFlow.WithContext(dbcache.WithCtx(c.ctx)).Where(qFlow.UUID.Eq(uuid)).UpdateSimple(updateCol)
	if err != nil {
		g.Log().Error(c.ctx, "update sign-off form fail:", err.Error())
	}
}

// 處理跑完所有流程後操作
func (c *checkIn) handleSignDone(form *types.CheckInRequestForm) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm = tx.CheckInRequestForm
			err   error
		)

		// 更新表單狀態
		_, err = qForm.WithContext(dbcache.WithCtx(c.ctx)).Where(qForm.ID.Eq(form.ID)).UpdateSimple(qForm.SignStatus.Value(enum.SignStatusApprove))
		if err != nil {
			return err
		}

		// TODO: 寫入出勤狀態表

		// commit
		return nil
	})
}

// 寄信通知
func (c *checkIn) sendEmail(option *model.EmailOption[types.CheckInRequestForm]) {
	var (
		qEmployee     = query.Employee
		employeeEmail string
		url           string
		body          string
		err           error
	)

	// 查詢員工資訊
	err = qEmployee.WithContext(c.ctx).Select(qEmployee.Email).Where(qEmployee.ID.Eq(option.Form.EmployeeID)).Scan(&employeeEmail)
	if err != nil {
		option.Err = gberror.Newf("query employee email error:%s", err.Error())
		return
	}

	// get url
	url, err = utils.GetURL("/sign-off/checkIn")
	if err != nil {
		option.Err = gberror.Newf("get url error:%s", err.Error())
		return
	}

	url = fmt.Sprintf("%s?uuid=%s&locale=%s", url, option.UUID, option.Form.SignOffFlow[0].Locale)

	// template
	view := g.View("checkSign")
	body, err = view.ParseDefault(c.ctx, g.Map{
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
