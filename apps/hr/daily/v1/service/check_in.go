package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gbi18n "ghostbb.io/gb/i18n/gb_i18n"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbfile "ghostbb.io/gb/os/gb_file"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/daily/v1/model"
	"hrapi/apps/hr/sign-off/v1/service"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/paginator"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

func CheckIn(ctx context.Context, page ...*paginator.Pagination) ICheckIn {
	if len(page) != 0 {
		return &checkIn{ctx: ctx, page: page[0]}
	}
	return &checkIn{ctx: ctx}
}

type (
	ICheckIn interface {
		// GetCheckInFormByEmployee 查詢補打卡單
		GetCheckInFormByEmployee(in model.GetCheckInFormByEmployeeReq) (out model.GetCheckInFormByEmployeeRes, err error)
		// SaveCheckInForm 儲存補打卡單
		SaveCheckInForm(in model.SaveCheckInFormReq) (out model.SaveCheckInFormRes, err error)
		// DeleteCheckInForm 刪除補打卡單
		DeleteCheckInForm(in model.DeleteCheckInFormReq) error
	}

	checkIn struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetCheckInFormByEmployee 查詢補打卡單
func (c *checkIn) GetCheckInFormByEmployee(in model.GetCheckInFormByEmployeeReq) (out model.GetCheckInFormByEmployeeRes, err error) {
	var (
		qForm = query.CheckInRequestForm
		start = gbconv.Time(in.StartDate)
		end   = gbconv.Time(in.EndDate)
		conds = []gen.Condition{qForm.EmployeeID.Eq(in.EmployeeID)}
	)

	if in.StartDate != "" && in.EndDate != "" {
		conds = append(conds, qForm.CreatedAt.Between(start, end.AddDate(0, 0, 1)))
	}

	out.Items, err = qForm.WithContext(dbcache.WithCtx(c.ctx)).
		Scopes(c.page.Where(conds...)).
		Preload(field.Associations).
		Order(qForm.CreatedAt.Desc()).Find()
	if err != nil {
		return
	}

	out.Total = c.page.Total

	return
}

// SaveCheckInForm 儲存補打卡單
func (c *checkIn) SaveCheckInForm(in model.SaveCheckInFormReq) (out model.SaveCheckInFormRes, err error) {
	var (
		qForm = query.CheckInRequestForm
		form  = new(types.CheckInRequestForm)
	)

	if in.ID != 0 {
		// 如果id不為0代表進行更新
		// 查詢原有資料
		form, err = qForm.WithContext(c.ctx).Where(qForm.ID.Eq(in.ID)).First()
		if err != nil {
			return
		}
	}

	// 複製form(request body)到database model
	if err = copier.Copy(form, &in); err != nil {
		return
	}

	// 寫入基本訊息, 部門id, 代理人部門id等...
	if err = c.writeBasicInfo(form); err != nil {
		return
	}

	// 開啟事務
	err = query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm   = tx.CheckInRequestForm
			qFlow   = tx.CheckInSignOffFlow
			qDetail = tx.CheckInRequestFormDetail
			err     error
		)

		if form.ID == 0 {
			// 如果form id為0
			// create a leave form
			if err = qForm.WithContext(dbcache.WithCtx(c.ctx)).Create(form); err != nil {
				return err
			}
			// create order
			// maybe A20240131000010
			// A + 今天日期 + (id % 1000000)補全6位數
			form.Order = "C" + time.Now().Format("20060102") + fmt.Sprintf("%06d", form.ID%1000000)
		}

		// 刪除原本的detail
		_, err = qDetail.WithContext(dbcache.WithCtx(c.ctx)).Where(qDetail.CheckInRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除原本的簽核流程
		_, err = qFlow.WithContext(dbcache.WithCtx(c.ctx)).Unscoped().Where(qFlow.CheckInRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 創建流程
		if err = c.createFlow(form); err != nil {
			return err
		}

		// 附件檔案存檔
		var paths []string
		paths, err = c.attachSave(form.Order, in.AttachArray)
		if err != nil {
			return err
		}
		form.Attach = gbstr.Join(paths, ",")

		// 儲存至資料庫
		// create or update
		if err = qForm.WithContext(dbcache.WithCtx(c.ctx)).Save(form); err != nil {
			return err
		}

		// commit
		return nil
	})
	if err != nil {
		return
	}

	// 開啟一個新的service
	// 因為如果response的話，context會被cancel
	var locale = gbi18n.LanguageFromCtx(c.ctx)
	go service.CheckIn(gbi18n.WithLanguage(gbctx.New(), locale)).NextFlow(form.SignOffFlow[0].UUID)

	out.ID = form.ID
	out.Order = form.Order
	return out, nil
}

// DeleteCheckInForm 刪除補打卡單
func (c *checkIn) DeleteCheckInForm(in model.DeleteCheckInFormReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qForm          = tx.CheckInRequestForm
			qDetail        = tx.CheckInRequestFormDetail
			qFlow          = tx.CheckInSignOffFlow
			qCheckInStatus = tx.CheckInStatus
			empID          uint
			dateRange      struct {
				StartDate driver.Date
				EndDate   driver.Date
			}
			err error
		)

		// 搜尋此表單的員工
		err = qForm.WithContext(c.ctx).Where(qForm.ID.Eq(in.ID)).Pluck(qForm.EmployeeID, &empID)
		if err != nil {
			return err
		}

		// 查詢此表單最小日期和最大日期
		err = qDetail.WithContext(c.ctx).Select(qDetail.Date.Min().As("start_date"), qDetail.Date.Max().As("end_date")).
			Where(qDetail.CheckInRequestFormID.Eq(in.ID)).Scan(&dateRange)
		if err != nil {
			return err
		}

		// 刪除此表單
		_, err = qForm.WithContext(dbcache.WithCtx(c.ctx)).Where(qForm.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除明細
		_, err = qDetail.WithContext(dbcache.WithCtx(c.ctx)).Where(qDetail.CheckInRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除流程
		_, err = qFlow.WithContext(dbcache.WithCtx(c.ctx)).Where(qFlow.CheckInRequestFormID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 更新出勤狀態表
		err = qCheckInStatus.WithContext(dbcache.WithCtx(c.ctx)).
			UpdateStatus(
				dateRange.StartDate.AddDate(0, 0, -1).Format(),
				dateRange.EndDate.AddDate(0, 0, 1).Format(),
				gbconv.String(empID),
			)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// 寫入基本訊息
func (c *checkIn) writeBasicInfo(form *types.CheckInRequestForm) (err error) {
	var (
		qEmployee = query.Employee
	)

	// 查詢員工部門id
	err = qEmployee.WithContext(dbcache.WithCtx(c.ctx)).Select(qEmployee.DepartmentId).Where(qEmployee.ID.Eq(form.EmployeeID)).Scan(&form.DepartmentID)
	if err != nil {
		return
	}

	return nil
}

// 創建流程
func (c *checkIn) createFlow(form *types.CheckInRequestForm) error {
	var (
		qFlowSetting = query.CheckInSignOffSetting
		qDepartment  = query.Department
		flowSetting  []*types.CheckInSignOffSetting
		err          error
	)

	// 查詢簽核設定
	flowSetting, err = qFlowSetting.WithContext(dbcache.WithCtx(c.ctx)).Find()
	if err != nil {
		return err
	}

	// 查詢部門所有父子關係
	// with tree as ((?) union all select b.* from tree inner join department b on b.id = tree.parent_id
	var departments []*types.Department
	departments, err = qDepartment.WithContext(c.ctx).QueryUp(form.DepartmentID)
	if err != nil {
		return err
	}

	// 流程判斷
	var deptNowID = form.DepartmentID // 記錄當前部門id, 方便向上查詢

	for _, fs := range flowSetting {
		var signEmpID = uint(0) // 紀錄當前關卡需簽核的員工

		// 如果條件(>=天數)為0或是達成條件
		// 0代表無論如何都必須使用此關卡
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

		// 不等於0代表條件達成，加入流程
		// 等於0代表向上沒有部門或其他邏輯問題
		if signEmpID != 0 {
			form.SignOffFlow = append(form.SignOffFlow, &types.CheckInSignOffFlow{
				CheckInRequestFormID: form.ID,
				SignOffEmployeeID:    signEmpID,
				SignType:             fs.SignType,
				Notify:               fs.Notify,
				Remark:               fs.Remark,
			})
		}
	}

	// 最後一關固定為通知起單人
	// 簽核完成後通知起單人
	form.SignOffFlow = append(form.SignOffFlow, &types.CheckInSignOffFlow{
		CheckInRequestFormID: form.ID,
		SignOffEmployeeID:    form.EmployeeID,
		SignType:             enum.SpecificEmployee, // 特定人員:起單人
		Notify:               enum.NotifyOnly,       // 僅通知
	})

	// 遍歷關卡
	for index, flow := range form.SignOffFlow {
		flow.Level = uint(index + 1)                             // 重新定義關卡
		flow.UUID = uuid.Must(uuid.NewRandom()).String()         // 為每一關生成uuid
		flow.Locale = enum.Locale(gbi18n.LanguageFromCtx(c.ctx)) // 語言為請求時使用的語言
	}

	return nil
}

// 附件儲存
func (c *checkIn) attachSave(order string, paths []string) ([]string, error) {
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
