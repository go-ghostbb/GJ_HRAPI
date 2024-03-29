package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/salary/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/paginator"
	"math"
)

func Salary(ctx context.Context, page ...*paginator.Pagination) ISalary {
	if len(page) != 0 {
		return &salary{ctx: ctx, page: page[0]}
	}
	return &salary{ctx: ctx}
}

type (
	ISalary interface {
		// Calc 薪資作業
		Calc(in model.SalaryCalcReq) (err error)
		// GetSalaryCalc 查詢所有薪資作業
		GetSalaryCalc() (out model.GetSalaryCalcRes, err error)
		// 刪除薪資作業
		Delete(in model.DeleteSalaryCalcReq) error
		// getByIDSalaryCalc 根據ID查詢薪資作業
		GetByIDSalaryCalc(in model.GetByIDSalaryCalcReq) (out model.GetByIDSalaryCalcRes, err error)
		// SaveCalcEmployee 儲存設定
		SaveCalcEmployee(in model.SaveEmployeeItemReq) error
	}

	salary struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// Calc 薪資作業
func (s *salary) Calc(in model.SalaryCalcReq) (err error) {
	var (
		qCalcSalary = query.CalcSalary
		result      *types.CalcSalary
	)

	// 建立基本訊息
	result, err = s.basicInfo(in)
	if err != nil {
		return
	}

	// 插入加項
	err = s.appendAddItem(in, result)
	if err != nil {
		return
	}

	// 插入減項
	err = s.appendReduceItem(in, result)
	if err != nil {
		return
	}

	return qCalcSalary.WithContext(dbcache.WithCtx(s.ctx)).Create(result)
}

// 填入基本訊息
func (s *salary) basicInfo(in model.SalaryCalcReq) (calc *types.CalcSalary, err error) {
	var (
		qEmployee      = query.Employee
		qConfig        = query.ConfigMap
		qCheckInStatus = query.CheckInStatus
		daysAMonth     float32
		hoursADay      float32
		startDate      = gbconv.Time(in.DateRangeStart)
		endDate        = gbconv.Time(in.DateRangeEnd)
		employees      []*types.Employee
	)
	employees, err = qEmployee.WithContext(s.ctx).Where(qEmployee.ID.In(in.EmployeeID...)).Find()
	if err != nil {
		return nil, err
	}

	// 尋找設定, 一個月算幾天
	err = qConfig.WithContext(s.ctx).Select(qConfig.Value).Where(qConfig.Key.Eq("DaysAMonth")).Scan(&daysAMonth)
	if err != nil {
		return nil, err
	}

	// 尋找設定, 一天幾小時
	err = qConfig.WithContext(s.ctx).Select(qConfig.Value).Where(qConfig.Key.Eq("HoursADay")).Scan(&hoursADay)
	if err != nil {
		return nil, err
	}

	calc = new(types.CalcSalary)
	calc.Start = startDate
	calc.End = endDate
	calc.Stage = 1
	calc.FounderEmployeeID = in.FounderEmployeeID
	for _, employee := range employees {
		temp := &types.CalcSalaryEmployee{
			EmployeeID: employee.ID,
			Employee:   employee,
			CalcSalary: calc,
		}

		switch employee.SalaryCycle {
		case enum.SalaryMonth:
			// 月薪制
			temp.Salary = employee.Salary
			// 將月薪轉換成時薪
			temp.HourlySalary = employee.Salary / daysAMonth / hoursADay
		case enum.SalaryHour:
			// 時薪制
			temp.HourlySalary = employee.Salary
			// 計算總上班時數
			totalHours, err := qCheckInStatus.WithContext(s.ctx).QueryTotalAttendHours(employee.ID, in.DateRangeStart, in.DateRangeEnd)
			if err != nil {
				return nil, err
			}
			temp.Salary = employee.Salary * float32(math.Floor(float64(totalHours)))
		}

		calc.CalcSalaryEmployee = append(calc.CalcSalaryEmployee, temp)
	}

	return calc, nil
}

// 套用加項
func (s *salary) appendAddItem(in model.SalaryCalcReq, calc *types.CalcSalary) (err error) {
	var (
		qAddItem = query.SalaryAddItem
		addItems []*struct {
			EmployeeID      uint
			SalaryAddItemID uint
			SalaryType      enum.SalaryType
			IncomeTax       bool
			Benefits        bool
			Premiums        bool
			Amount          float32
			Multiply        float32
		}
		queryRes     []g.Map
		empIDCalcMap = make(map[uint]*types.CalcSalaryEmployee)
	)

	queryRes, err = qAddItem.WithContext(s.ctx).QueryByEmployeeID(gbstr.JoinAny(in.EmployeeID, ","), in.DateRangeStart, in.DateRangeEnd)
	if err != nil {
		return err
	}

	if err = gbconv.Structs(queryRes, &addItems); err != nil {
		return err
	}

	// 建立map對照
	for _, c := range calc.CalcSalaryEmployee {
		if _, ok := empIDCalcMap[c.EmployeeID]; !ok {
			empIDCalcMap[c.EmployeeID] = c
		}
	}

	// 遍歷
	for _, a := range addItems {
		if a.Multiply == 0 {
			continue
		}
		empIDCalcMap[a.EmployeeID].SalaryAdd = append(empIDCalcMap[a.EmployeeID].SalaryAdd, &types.CalcSalaryAdd{
			CalcSalaryEmployee: empIDCalcMap[a.EmployeeID],
			SalaryAddItemID:    a.SalaryAddItemID,
			Amount:             a.Amount * a.Multiply,
			SalaryType:         a.SalaryType,
			IncomeTax:          a.IncomeTax,
			Benefits:           a.Benefits,
			Premiums:           a.Premiums,
		})
	}

	return nil
}

// 套用減項
func (s *salary) appendReduceItem(in model.SalaryCalcReq, calc *types.CalcSalary) (err error) {
	var (
		qReduceItem = query.SalaryReduceItem
		reduceItems []*struct {
			EmployeeID         uint
			SalaryReduceItemID uint
			IncomeTax          bool
			Amount             float32
			Multiply           float32
		}
		queryRes     []g.Map
		empIDCalcMap = make(map[uint]*types.CalcSalaryEmployee)
	)

	queryRes, err = qReduceItem.WithContext(s.ctx).QueryByEmployeeID(gbstr.JoinAny(in.EmployeeID, ","), in.DateRangeStart, in.DateRangeEnd)
	if err != nil {
		return err
	}

	if err = gbconv.Structs(queryRes, &reduceItems); err != nil {
		return err
	}

	// 建立map對照
	for _, c := range calc.CalcSalaryEmployee {
		if _, ok := empIDCalcMap[c.EmployeeID]; !ok {
			empIDCalcMap[c.EmployeeID] = c
		}
	}

	// 遍歷
	for _, a := range reduceItems {
		if a.Multiply == 0 {
			continue
		}
		empIDCalcMap[a.EmployeeID].SalaryReduce = append(empIDCalcMap[a.EmployeeID].SalaryReduce, &types.CalcSalaryReduce{
			CalcSalaryEmployee: empIDCalcMap[a.EmployeeID],
			SalaryReduceItemID: a.SalaryReduceItemID,
			Amount:             a.Amount * a.Multiply,
			IncomeTax:          a.IncomeTax,
		})
	}

	return nil
}

// GetSalaryCalc 查詢所有薪資作業
func (s *salary) GetSalaryCalc() (out model.GetSalaryCalcRes, err error) {
	var (
		qCalc    = query.CalcSalary
		calcList []*types.CalcSalary
	)

	calcList, err = qCalc.WithContext(dbcache.WithCtx(s.ctx)).Preload(
		field.Associations,
		qCalc.CalcSalaryEmployee.Employee,
		qCalc.CalcSalaryEmployee.SalaryAdd,
		qCalc.CalcSalaryEmployee.SalaryAdd.SalaryAddItem,
		qCalc.CalcSalaryEmployee.SalaryReduce,
		qCalc.CalcSalaryEmployee.SalaryReduce.SalaryReduceItem,
	).Scopes(s.page.Where()).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, calcList); err != nil {
		return
	}

	out.Total = s.page.Total

	return out, nil
}

// Delete 刪除薪資作業
func (s *salary) Delete(in model.DeleteSalaryCalcReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qCalc         = tx.CalcSalary
			qCalcAdd      = tx.CalcSalaryAdd
			qCalcReduce   = tx.CalcSalaryReduce
			qCalcEmployee = tx.CalcSalaryEmployee
			subQueryCEID  = qCalcEmployee.WithContext(dbcache.WithCtx(s.ctx)).Select(qCalcEmployee.ID).Where(qCalcEmployee.CalcSalaryID.Eq(in.ID))
			err           error
		)

		// 刪除加項減項
		_, err = qCalcAdd.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalcAdd.Columns(qCalcAdd.CalcSalaryEmployeeID).In(subQueryCEID)).Delete()
		if err != nil {
			return err
		}
		_, err = qCalcReduce.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalcReduce.Columns(qCalcReduce.CalcSalaryEmployeeID).In(subQueryCEID)).Delete()
		if err != nil {
			return err
		}

		// 刪除計算的員工
		_, err = qCalcEmployee.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalcEmployee.CalcSalaryID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除薪資作業
		_, err = qCalc.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalc.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// getByIDSalaryCalc 根據ID查詢薪資作業
func (s *salary) GetByIDSalaryCalc(in model.GetByIDSalaryCalcReq) (out model.GetByIDSalaryCalcRes, err error) {
	var (
		qCalc = query.CalcSalary
	)

	out.CalcSalary, err = qCalc.WithContext(dbcache.WithCtx(s.ctx)).Preload(
		field.Associations,
		qCalc.CalcSalaryEmployee.Employee,
		qCalc.CalcSalaryEmployee.SalaryAdd,
		qCalc.CalcSalaryEmployee.SalaryAdd.SalaryAddItem,
		qCalc.CalcSalaryEmployee.SalaryReduce,
		qCalc.CalcSalaryEmployee.SalaryReduce.SalaryReduceItem,
	).Where(qCalc.ID.Eq(in.ID)).First()

	return
}

// SaveCalcEmployee 儲存設定
func (s *salary) SaveCalcEmployee(in model.SaveEmployeeItemReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qCalcEmployee = tx.CalcSalaryEmployee
			qCalcAdd      = tx.CalcSalaryAdd
			qCalcReduce   = tx.CalcSalaryReduce
			addItem       = make([]*types.CalcSalaryAdd, 0)
			reduceItem    = make([]*types.CalcSalaryReduce, 0)
			allIDs        = make([]uint, 0) // 搜集所有id
			err           error
		)

		for _, calcEmployee := range in.CalcSalaryEmployee {
			// 尋找所有加項
			for _, item := range calcEmployee.SalaryAdd {
				addItem = append(addItem, item)
			}

			// 尋找所有減項
			for _, item := range calcEmployee.SalaryReduce {
				reduceItem = append(reduceItem, item)
			}

			allIDs = append(allIDs, calcEmployee.ID)
		}

		// 刪除原有加項
		_, err = qCalcAdd.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalcAdd.CalcSalaryEmployeeID.In(allIDs...)).Unscoped().Delete()
		if err != nil {
			return err
		}
		// 新增新的加項
		err = qCalcAdd.WithContext(dbcache.WithCtx(s.ctx)).Create(addItem...)
		if err != nil {
			return err
		}

		// 儲存減項
		_, err = qCalcReduce.WithContext(dbcache.WithCtx(s.ctx)).Where(qCalcReduce.CalcSalaryEmployeeID.In(allIDs...)).Unscoped().Delete()
		if err != nil {
			return err
		}
		err = qCalcReduce.WithContext(dbcache.WithCtx(s.ctx)).Create(reduceItem...)
		if err != nil {
			return err
		}

		err = qCalcEmployee.WithContext(dbcache.WithCtx(s.ctx)).Save(in.CalcSalaryEmployee...)
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}
