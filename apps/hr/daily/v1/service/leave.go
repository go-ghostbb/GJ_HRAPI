package service

import (
	"context"
	"database/sql"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbi18n "ghostbb.io/gb/i18n/gb_i18n"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbfile "ghostbb.io/gb/os/gb_file"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"hrapi/apps/hr/daily/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/email"
	"hrapi/internal/utils/identity"
	"hrapi/internal/utils/paginator"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func Leave(ctx context.Context, page ...*paginator.Pagination) ILeave {
	if len(page) != 0 {
		return &leave{ctx: ctx, page: page[0]}
	}
	return &leave{ctx: ctx}
}

type (
	ILeave interface {
		// GetLeaveFormByEmployee 獲取請假單列表
		GetLeaveFormByEmployee(in model.GetLeaveFormByEmployeeReq) (out model.GetLeaveFormByEmployeeRes, err error)
		// DeleteLeaveForm 刪除請假單
		DeleteLeaveForm(in model.DeleteLeaveFormReq) error
		// SaveLeaveForm is to insert or update leave request form
		SaveLeaveForm(in model.SaveLeaveFormReq) (out model.SaveLeaveFormRes, err error)
		// 寄信通知
		sendEmail(form *types.LeaveRequestForm)
	}

	leave struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetLeaveFormByEmployee 獲取請假單列表
func (l *leave) GetLeaveFormByEmployee(in model.GetLeaveFormByEmployeeReq) (out model.GetLeaveFormByEmployeeRes, err error) {
	var (
		qForm = query.LeaveRequestForm
		start = gbconv.Time(in.StartDate)
		end   = gbconv.Time(in.EndDate)
		conds = []gen.Condition{qForm.EmployeeID.Eq(in.EmployeeID)}
	)

	if in.StartDate != "" && in.EndDate != "" {
		conds = append(conds, qForm.CreatedAt.Between(start, end.AddDate(0, 0, 1)))
	}

	out.Items, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).
		Scopes(l.page.Where(conds...)).
		Preload(field.Associations).
		Order(qForm.CreatedAt.Desc()).Find()
	if err != nil {
		return
	}

	out.Total = l.page.Total

	return
}

// DeleteLeaveForm 刪除請假單
func (l *leave) DeleteLeaveForm(in model.DeleteLeaveFormReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm          = tx.LeaveRequestForm
			qFlow          = tx.LeaveSignOffFlow
			qCheckInStatus = tx.CheckInStatus
			qLeaveCorrect  = tx.LeaveCorrect
			form           *types.LeaveRequestForm
			err            error
		)

		// 查詢此表單內容
		form, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).Where(qForm.ID.Eq(in.ID)).First()
		if err != nil {
			return err
		}

		// 刪除此表單
		_, err = qForm.WithContext(dbcache.WithCtx(l.ctx)).Where(qForm.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除流程
		_, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.LeaveRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 更新請假可用表
		_, err = qLeaveCorrect.WithContext(l.ctx).UpdateCount(form.ID, false)
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

// SaveLeaveForm is to insert or update leave request form
func (l *leave) SaveLeaveForm(in model.SaveLeaveFormReq) (out model.SaveLeaveFormRes, err error) {
	var (
		qForm     = query.LeaveRequestForm
		leaveForm = new(types.LeaveRequestForm)
	)

	if in.ID != 0 {
		// 如果id不為0代表進行更新
		// 查詢原有資料
		leaveForm, err = qForm.WithContext(l.ctx).Where(qForm.ID.Eq(in.ID)).First()
		if err != nil {
			return
		}
	}

	// 複製form(request body)到database model
	if err = copier.Copy(leaveForm, &in); err != nil {
		return
	}

	// 判斷時間有沒有重疊
	var overlapCheck bool
	if overlapCheck, err = l.isOverlap(leaveForm); err != nil {
		return
	} else if overlapCheck {
		// 回傳時間重疊錯誤
		err = gberror.New("time overlap")
		return
	}

	// 寫入基本訊息, 部門id, 代理人部門id等...
	if err = l.writeBasicInfo(leaveForm); err != nil {
		return
	}

	// 開始資料庫操作，開啟事務
	opts := &sql.TxOptions{
		// 啟用讀取未提交
		Isolation: sql.LevelReadCommitted,
	}
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm          = tx.LeaveRequestForm
			qFlow          = tx.LeaveSignOffFlow
			qCheckInStatus = tx.CheckInStatus
			qLeaveCorrect  = tx.LeaveCorrect
			err            error
		)

		// 如果form id為0，先建立資料獲取id
		if leaveForm.ID == 0 {
			leaveForm.Order, err = identity.LeaveOrder(l.ctx)
			if err != nil {
				return err
			}
		}

		// 刪除該form的簽核流程並且建立新的
		// delete flow
		if _, err = qFlow.WithContext(dbcache.WithCtx(l.ctx)).Where(qFlow.LeaveRequestFormID.Eq(leaveForm.ID)).Unscoped().Delete(); err != nil {
			return err
		}
		// create flow
		if err = l.createFlow(leaveForm); err != nil {
			return err
		}
		leaveForm.SignStatus = enum.SignStatusSending // 送簽中

		// 附件檔案存檔
		var paths []string
		paths, err = l.attachSave(leaveForm.Order, in.AttachArray)
		if err != nil {
			return err
		}
		leaveForm.Attach = gbstr.Join(paths, ",")

		// 儲存至資料庫
		// create or update
		if err = qForm.WithContext(dbcache.WithCtx(l.ctx)).Save(leaveForm); err != nil {
			return err
		}

		// 更新請假可用表
		_, err = qLeaveCorrect.WithContext(l.ctx).UpdateCount(leaveForm.ID, true)
		if err != nil {
			return err
		}

		// 更新出勤狀態表
		err = qCheckInStatus.WithContext(dbcache.WithCtx(l.ctx)).
			UpdateStatus(
				leaveForm.StartDate.AddDate(0, 0, -1).Format(),
				leaveForm.EndDate.AddDate(0, 0, 1).Format(),
				gbconv.String(leaveForm.EmployeeID),
			)
		if err != nil {
			return err
		}

		// commit
		return nil
	}, opts)
	if err != nil {
		return
	}

	// send email
	// 開啟一個新的service
	// 因為如果response的話，context會被cancel
	var locale = gbi18n.LanguageFromCtx(l.ctx)
	go Leave(gbi18n.WithLanguage(gbctx.New(), locale)).sendEmail(leaveForm)

	// return
	out.ID = leaveForm.ID
	out.Order = leaveForm.Order
	return out, nil
}

// 寫入基本訊息
func (l *leave) writeBasicInfo(form *types.LeaveRequestForm) (err error) {
	var (
		qEmployee = query.Employee
	)

	// 查詢員工部門id
	err = qEmployee.WithContext(dbcache.WithCtx(l.ctx)).Select(qEmployee.DepartmentId).Where(qEmployee.ID.Eq(form.EmployeeID)).Scan(&form.DepartmentID)
	if err != nil {
		return
	}

	// 查詢代理員工部門id
	err = qEmployee.WithContext(dbcache.WithCtx(l.ctx)).Select(qEmployee.DepartmentId).Where(qEmployee.ID.Eq(form.ProxyEmployeeID)).Scan(&form.ProxyDepartmentID)
	if err != nil {
		return
	}

	// 計算請假時間
	err = l.countTime(form)
	if err != nil {
		return
	}

	return nil
}

// 計算請假時間
func (l *leave) countTime(form *types.LeaveRequestForm) (err error) {
	var (
		qConfigMap    = query.ConfigMap
		qWorkSchedule = query.WorkSchedule
		qLeave        = query.Leave
		hoursADay     float32
	)

	form.Leave, err = qLeave.WithContext(dbcache.WithCtx(l.ctx)).Where(qLeave.ID.Eq(form.LeaveID)).First()
	if err != nil {
		if gberror.Is(err, gorm.ErrRecordNotFound) {
			return gberror.Newf("leave id:%d not found", form.LeaveID)
		}
		return err
	}

	// 取得設定檔中一天算幾小時
	err = qConfigMap.WithContext(l.ctx).Select(qConfigMap.Value).Where(qConfigMap.Key.Eq("HoursADay")).Scan(&hoursADay)
	if err != nil {
		if gberror.Is(err, gorm.ErrRecordNotFound) {
			g.Log().Errorf(l.ctx, `config map not found: "%s"`, "HoursADay")
		}
		return err
	}

	// 班表遍歷+計算
	var (
		minute float32
		hour   float32
		day    float32
	)

	tmpStart := form.StartDate
	for tmpStart.Before(form.EndDate.AddDate(0, 0, 1)) {
		start := tmpStart.DateTime(form.StartTime)
		end := tmpStart.DateTime(form.EndTime)
		if form.StartTime.After(form.EndTime) {
			// 如果開始時間>結束時間，代表過夜
			end = end.AddDate(0, 0, 1)
		}

		var tmpHour float32
		tmpHour, err = qWorkSchedule.WithContext(l.ctx).WorkHourCount(form.EmployeeID, start.Format(), end.Format())
		if err != nil {
			return err
		}
		hour += tmpHour

		// 判斷最小單位
		if form.Leave.Minimum != 0 {
			temp, _ := decimal.NewFromFloat32(hour * 60).
				Div(decimal.NewFromInt(int64(form.Leave.Minimum))).Ceil().
				Mul(decimal.NewFromInt(int64(form.Leave.Minimum))).
				Float64()
			hour = float32(temp / 60)
		}

		// step
		tmpStart = tmpStart.AddDate(0, 0, 1)
	}

	minute = hour * 60
	day = hour / hoursADay

	// 請假時數不能為0
	if hour == 0 {
		return gberror.New("the number of leave hours cannot be 0")
	}

	// 回寫計算結果
	form.LeaveDayCount = day
	form.LeaveHourCount = hour
	form.LeaveMinuteCount = minute

	return nil
}

// 判斷時間是否有重疊
func (l *leave) isOverlap(form *types.LeaveRequestForm) (bool, error) {
	var (
		qForm = query.LeaveRequestForm
	)

	// 查詢請假單中有重疊的日期
	queryRes, err := qForm.WithContext(l.ctx).Where(
		qForm.ID.Neq(form.ID),
		// isOverlap: A1 <= B2 and A2 >= B1
		qForm.StartDate.Lte(form.EndDate), qForm.EndDate.Gte(form.StartDate),
		qForm.SignStatus.In(enum.SignStatusSending, enum.SignStatusApprove, enum.SignStatusSigning),
		qForm.EmployeeID.Eq(form.EmployeeID),
	).Find()
	if err != nil {
		return false, err
	}

	// 判斷重疊的時間裡時間是不是也重疊
	// 遍歷尋出來的表單
	for _, f := range queryRes {
		// 遍歷表單(f)日期
		for f.EndDate.AddDate(0, 0, 1).After(f.StartDate) {
			// A1 < B2 && A2 > B1
			var (
				A1 = gbconv.Time(f.StartDate.Format() + " " + f.StartTime.Format())
				A2 = gbconv.Time(f.StartDate.Format() + " " + f.EndTime.Format())
			)

			// 如果開始時間>結束時間
			// 代表隔夜，需要+1天
			if f.StartTime.Unix() > f.EndTime.Unix() {
				A2 = A2.AddDate(0, 0, 1)
			}

			tempStartDate := form.StartDate
			// 遍歷表單(form)日期
			for form.EndDate.AddDate(0, 0, 1).After(tempStartDate) {
				var (
					B1 = gbconv.Time(form.StartDate.Format() + " " + form.StartTime.Format())
					B2 = gbconv.Time(form.StartDate.Format() + " " + form.EndTime.Format())
				)

				// 如果開始時間>結束時間
				// 代表隔夜，需要+1天
				if form.StartTime.Unix() > form.EndTime.Unix() {
					B2 = B2.AddDate(0, 0, 1)
				}

				// 判斷是否重疊
				// 公式：A1 < B2 && A2 > B1
				if A1.Unix() < B2.Unix() && A2.Unix() > B1.Unix() {
					return true, nil
				}

				// 迴圈step(1 day)
				tempStartDate = tempStartDate.AddDate(0, 0, 1)
			}

			// 迴圈step(1 day)
			f.StartDate = f.StartDate.AddDate(0, 0, 1)
		}
	}

	return false, nil
}

// 創建流程
func (l *leave) createFlow(form *types.LeaveRequestForm) error {
	var (
		qFlowSetting = query.LeaveSignOffSetting
		qDepartment  = query.Department
		flowSetting  []*types.LeaveSignOffSetting
		err          error
	)

	// 查詢請假簽核設定
	flowSetting, err = qFlowSetting.WithContext(l.ctx).
		Where(qFlowSetting.DepartmentID.Eq(form.DepartmentID), qFlowSetting.LeaveID.Eq(form.LeaveID)).
		Order(qFlowSetting.Level).Find()
	if err != nil {
		return err
	}

	// 查詢部門所有父子關係
	// with tree as ((?) union all select b.* from tree inner join department b on b.id = tree.parent_id
	var departments []*types.Department
	departments, err = qDepartment.WithContext(l.ctx).QueryUp(form.DepartmentID)
	if err != nil {
		return err
	}

	// 簽核流程中，第一關一定為代理人，直接加入form簽核流程中
	form.SignOffFlow = append(form.SignOffFlow, &types.LeaveSignOffFlow{
		LeaveRequestFormID: form.ID,
		SignOffEmployeeID:  form.ProxyEmployeeID,   // 代理員工ID
		SignType:           enum.Agent,             // 代理人
		Status:             enum.SignStatusSigning, // 簽核中
		Notify:             enum.SignOffPlusNotify, // 簽核通知
	})

	// 流程判斷
	var deptNowID = form.DepartmentID // 記錄當前部門id, 方便向上查詢

	for _, fs := range flowSetting {
		var signEmpID = uint(0) // 紀錄當前關卡需簽核的員工

		// 如果條件(>=天數)為0或是達成條件
		// 0代表無論如何都必須使用此關卡
		if fs.GteDay == 0 || form.LeaveDayCount >= fs.GteDay {
			switch fs.SignType {
			case enum.DepartmentManager:
				// 簽核對象為部門主管
				for {
					index := slices.IndexFunc(departments, func(d *types.Department) bool {
						// 找到當前部門
						return d.ID == deptNowID
					})

					// 如果不是-1代表有找到部門
					if index != -1 {
						if departments[index].ManagerId == form.EmployeeID {
							// 如果部門主管為起單人，繼續向上搜尋
							deptNowID = departments[index].ParentID
							continue
						} else {
							// 如果不是起單人，代表此關卡簽核員工為此部門的主管
							signEmpID = departments[index].ManagerId
						}
					}

					// 如果沒搜尋到或是搜尋到且部門主管不是起單人則離開
					break
				}
			case enum.SpecificEmployee:
				// 簽核對象為特定人員
				// 如果特定人員等於起單人，則不加入流程
				if fs.SpecificEmployeeID != form.EmployeeID {
					signEmpID = fs.SpecificEmployeeID
				}
			}
		}

		// 不等於0代表條件達成，加入流程
		// 等於0代表向上沒有部門或其他邏輯問題
		if signEmpID != 0 {
			form.SignOffFlow = append(form.SignOffFlow, &types.LeaveSignOffFlow{
				LeaveRequestFormID: form.ID,
				SignOffEmployeeID:  signEmpID,
				SignType:           fs.SignType,
				Notify:             fs.Notify,
				Remark:             fs.Remark,
			})
		}
	}

	// 最後一關固定為通知起單人
	// 簽核完成後通知起單人
	form.SignOffFlow = append(form.SignOffFlow, &types.LeaveSignOffFlow{
		LeaveRequestFormID: form.ID,
		SignOffEmployeeID:  form.EmployeeID,
		SignType:           enum.SpecificEmployee, // 特定人員:起單人
		Notify:             enum.NotifyOnly,       // 僅通知
	})

	// 遍歷關卡
	for index, flow := range form.SignOffFlow {
		flow.Level = uint(index + 1)                             // 重新定義關卡
		flow.UUID = uuid.Must(uuid.NewRandom()).String()         // 為每一關生成uuid
		flow.Locale = enum.Locale(gbi18n.LanguageFromCtx(l.ctx)) // 語言為請求時使用的語言
	}

	return nil
}

// 附件儲存
func (l *leave) attachSave(order string, paths []string) ([]string, error) {
	if len(paths) == 0 {
		return paths, nil
	}

	var (
		result     = make([]string, 0)
		dirBuilder strings.Builder
	)

	dirBuilder.Write([]byte("assets"))
	if order != "" {
		dirBuilder.Write([]byte("/" + order)) // assets/A20231206XXXXXX
	}

	// maybe assets/tmp/...
	for _, path := range paths {
		if !gbstr.Contains(path, "assets/tmp") {
			// 如果該文件不再暫存區，直接加入結果
			result = append(result, path)
			continue
		}

		// 開啟暫存區的文件
		src, err := os.Open(path)
		if err != nil {
			if src != nil {
				_ = src.Close()
			}
			return nil, err
		}

		// mkdir dir
		utils.MkdirIfNotExist(dirBuilder.String())

		// maybe assets/A20231206XXXXXX/filename
		var filePath = filepath.Join(dirBuilder.String(), gbfile.Basename(path))
		out, err := os.Create(filePath)
		if err != nil {
			if out != nil {
				_ = out.Close()
			}
			return nil, err
		}

		// 將暫存區檔案複製到此請假單靜態文件資料夾區
		_, err = io.Copy(out, src)
		if err != nil {
			if src != nil {
				_ = src.Close()
			}
			if out != nil {
				_ = out.Close()
			}
			return nil, err
		}

		filePath = gbstr.Replace(filePath, `\`, "/")
		result = append(result, filePath)

		// close
		_ = src.Close()
		_ = out.Close()
	}

	return result, nil
}

// 寄信通知
func (l *leave) sendEmail(form *types.LeaveRequestForm) {
	var (
		qEmployee     = query.Employee
		employeeEmail string
		url           string
		body          string
		err           error
	)

	// 查詢員工資訊
	err = qEmployee.WithContext(l.ctx).Select(qEmployee.Email).Where(qEmployee.ID.Eq(form.EmployeeID)).Scan(&employeeEmail)
	if err != nil {
		g.Log().Error(l.ctx, "query employee email error:", err.Error())
		return
	}

	// get url
	url, err = utils.GetURL("/sign-off/leave")
	if err != nil {
		g.Log().Error(l.ctx, "get url error:", err.Error())
		return
	}

	url = fmt.Sprintf("%s?uuid=%s&locale=%s", url, form.SignOffFlow[0].UUID, form.SignOffFlow[0].Locale)

	// template
	view := g.View("leaveSign")
	body, err = view.ParseDefault(l.ctx, g.Map{
		"url":     url,
		"content": form.Order,
		"msg":     "請假單簽核",
	})
	if err != nil {
		g.Log().Error(l.ctx, "parse template(mail/leaveSign.html) error:", err.Error())
		return
	}

	// send
	err = email.Service.SendTo([]string{employeeEmail}, fmt.Sprintf("待簽核【請假單】- %s", form.Order), body)
	if err != nil {
		g.Log().Error(l.ctx, "send email error:", err.Error())
	}
}
