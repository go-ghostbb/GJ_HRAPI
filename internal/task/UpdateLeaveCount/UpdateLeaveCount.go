package UpdateLeaveCount

import (
	"context"
	"fmt"
	gberror "ghostbb.io/gb/errors/gb_error"
	gblog "ghostbb.io/gb/os/gb_log"
	"gorm.io/gorm"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
	"time"
)

type LeaveCheck struct {
	// Q
	// @Description: 資料庫對象
	Q *query.Query
	// Log
	// @Description: Log對象
	Log *gblog.Logger
	// seniority
	// @Description: 年資
	seniority map[employeeID]seniority
	// seniority
	// @Description: 年底年資
	endYear map[employeeID]seniority
	// core
	// @Description: 可用天數
	core map[employeeID]map[leaveID]coreInfo
	// now
	// @Description: 今天日期
	now time.Time
}

type coreInfo struct {
	// day
	// @Description: 可請天數
	day float64
	// cycle
	// @Description: 週期制度
	cycle enum.LeaveCycle
	// isDefer
	// @Description: 是否遞延
	isDefer bool
	// deferDay
	// @Description: 遞延最大天數
	deferDay uint
	// isMaturity
	// @Description: 是否自訂過期
	isMaturity bool
	// maturity
	// @Description: 遞延過期
	maturity uint
	// update
	// @Description: 是否更新
	update bool
	// currCond
	// @Description: 套用的條件
	currCond *types.LeaveGroupCondition
	// nextCond
	// @Description: 下個條件
	nextCond uint
	// nextDay
	// @Description: 下個條件給予天數
	nextDay uint
	// noGroup
	// @Description: 是否有群組
	noGroup bool
}

type employeeID uint

type leaveID uint

// seniority
// @Description: 年資訊息
type seniority struct {
	// hireDate
	// @Description: 入職時間
	hireDate time.Time
	// year
	// @Description: 年資(年)
	year int
	// month
	// @Description: 年資(月)
	month int
}

func (l *LeaveCheck) Init() {
	l.seniority = make(map[employeeID]seniority)
	l.endYear = make(map[employeeID]seniority)
	l.core = make(map[employeeID]map[leaveID]coreInfo)
	l.now, _ = time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
}

// Run
//
//	@Description: 定時任務入口
//	@receiver l
func (l *LeaveCheck) Run(ctx context.Context) {
	l.Init()

	var (
		qEmployee = l.Q.Employee
		qLeave    = l.Q.Leave
	)

	// 查詢年資
	employeeSeniority, err := qEmployee.WithContext(ctx).CalculateEmployeeSeniority()
	if err != nil {
		l.Log.Error(ctx, "calculate user seniority error", err)
		return
	}

	// 放入結構體
	l.seniority = make(map[employeeID]seniority) // 初始化
	for _, s := range employeeSeniority {
		l.seniority[employeeID(s["employee_id"].(int64))] = seniority{hireDate: s["hire_date"].(time.Time), year: int(s["year"].(int64)), month: int(s["month"].(int64))}
	}

	// 查詢年底年資
	endYear, err := qEmployee.WithContext(ctx).CalculateEmployeeSeniorityWithEndDate(fmt.Sprintf("%d-12-31", l.now.Year()))
	if err != nil {
		l.Log.Error(ctx, "calculate user seniority error", err)
		return
	}

	// 放入結構體
	l.endYear = make(map[employeeID]seniority) // 初始化
	for _, s := range endYear {
		l.endYear[employeeID(s["employee_id"].(int64))] = seniority{hireDate: s["hire_date"].(time.Time), year: int(s["year"].(int64)), month: int(s["month"].(int64))}
	}

	// 把所有假別取出
	leaves, err := qLeave.WithContext(ctx).
		Preload(qLeave.LeaveGroup).
		Preload(qLeave.LeaveGroup.Employee).
		Preload(qLeave.LeaveGroup.LeaveGroupCondition).
		Find()
	if err != nil {
		l.Log.Error(ctx, "query annual leave core error", err)
	}

	// 遍歷
	l.core = make(map[employeeID]map[leaveID]coreInfo) // 初始化
	for empID, sen := range l.seniority {
		l.core[empID] = make(map[leaveID]coreInfo) // 初始化
		for _, leave := range leaves {
			leaveID := leaveID(leave.ID)
			var group *types.LeaveGroup = nil

			// 查找被歸類在哪個群組
			for _, g := range leave.LeaveGroup {
				for _, e := range g.Employee {
					if uint(empID) == e.ID {
						group = g
						break
					}
				}
				// 如果有找到，退出迴圈
				if group != nil {
					break
				}
			}

			// 查找可用天數
			un := false
			if group != nil {
				// 先判斷第一關是否有達成條件
				if len(group.LeaveGroupCondition) >= 1 {
					switch group.LeaveGroupCondition[0].ConditionType {
					case enum.ConditionMonth:
						// 年資滿<#>個月
						if uint(sen.totalMonth()) >= group.LeaveGroupCondition[0].ConditionNum {
							un = true
						}
					case enum.ConditionYear:
						// 年資滿<#>年
						if uint(sen.year) >= group.LeaveGroupCondition[0].ConditionNum {
							un = true
						}
					}
				}

				if un {
					// 第一關有達成條件，才遍歷
					for index, c := range group.LeaveGroupCondition {
						core := coreInfo{
							day: float64(c.Result),
							// 未滿一年情況
							// 只有年資剛好滿足條件(月)和日期為(DD)同一天才會觸發更新
							update:     l.isUpdate(sen, c, leave.Cycle),
							cycle:      leave.Cycle,
							isDefer:    leave.Deferrable,
							deferDay:   leave.DeferrableDays,
							isMaturity: leave.CustomizableDuration,
							maturity:   leave.Duration,
							currCond:   c,
						}
						if index+1 < len(group.LeaveGroupCondition) {
							core.nextCond = group.LeaveGroupCondition[index+1].ConditionNum
							core.nextDay = group.LeaveGroupCondition[index+1].Result
						} else {
							core.nextCond = 0
							core.nextDay = 0
						}

						switch c.ConditionType {
						case enum.ConditionMonth:
							// 年資滿<#>個月
							if uint(sen.totalMonth()) >= c.ConditionNum {
								l.core[empID][leaveID] = core
							}
						case enum.ConditionYear:
							// 年資滿<#>年
							if uint(sen.year) >= c.ConditionNum {
								l.core[empID][leaveID] = core
							}
						}
					}
				}
			}

			// 如果找不到群組或是第一關條件未達成，帶入預設天數
			if group == nil || !un {
				core := coreInfo{
					day:        float64(leave.Default),
					update:     l.isUpdate(sen, nil, leave.Cycle),
					cycle:      leave.Cycle,
					isDefer:    leave.Deferrable,
					deferDay:   leave.DeferrableDays,
					isMaturity: leave.CustomizableDuration,
					maturity:   leave.Duration,
					noGroup:    group == nil,
				}
				if group != nil {
					core.currCond = group.LeaveGroupCondition[0]
					if len(group.LeaveGroupCondition) > 1 {
						core.nextCond = group.LeaveGroupCondition[1].ConditionNum
						core.nextDay = group.LeaveGroupCondition[1].Result
					}
				}
				l.core[empID][leaveID] = core
			}
		}
	}

	// 執行
	l.doUpdate(ctx)
}

// totalMonth
//
//	@Description: 計算總共多少個月
//	@receiver s
//	@return int
func (s *seniority) totalMonth() int {
	if s.year == 0 {
		return s.month
	}
	return s.year*12 + s.month
}

// isUpdate
//
//	@Description: 判斷是否更新
//	@receiver l
//	@param s
//	@param c
//	@param cycle
//	@return bool
func (l *LeaveCheck) isUpdate(s seniority, c *types.LeaveGroupCondition, cycle enum.LeaveCycle) bool {
	switch cycle {
	case enum.Annual:
		if c == nil {
			// 如果c = nil代表找不到群組或是第一關條件未達成
			// 第一關條件未達成代表為新員工，在建立新使用者時就會自動給假
			// 所以這邊只討論找不到群組狀況
			// 那更新時間就依入職時間
			return s.hireDate.Month() == l.now.Month() && s.hireDate.Day() == l.now.Day()
		}

		switch c.ConditionType {
		case enum.ConditionMonth:
			// 未滿一年情況
			// 只有年資剛好滿足條件(月)和日期為(DD)同一天才會觸發更新
			return uint(s.totalMonth()) == c.ConditionNum && s.hireDate.Day() == l.now.Day()
		case enum.ConditionYear:
			// 年資滿一年情況
			// 只有在月和天為同一天才會觸發更新
			return s.hireDate.Month() == l.now.Month() && s.hireDate.Day() == l.now.Day()
		default:
			return false
		}
	case enum.Calendar, enum.Default:
		// 曆年制和以「年」為單位給予都是每年一月一日更新
		return l.now.Month() == 1 && l.now.Day() == 1
	default:
		return false
	}
}

// doDefer
//
//	@Description: 遞延操作
//	@receiver l
//	@param tx
//	@param employeeID
//	@param leaveID
//	@param info
//	@return error
func (l *LeaveCheck) doDefer(ctx context.Context, tx *query.Query, employeeID employeeID, leaveID leaveID) error {
	var (
		qCorrect = tx.LeaveCorrect
		qDefer   = tx.LeaveDefer
	)
	left := 0.0
	// 查詢剩餘天數
	if err := qCorrect.WithContext(ctx).Select(qCorrect.Available.SubCol(qCorrect.Used)).Where(qCorrect.EmployeeID.Eq(uint(employeeID)), qCorrect.LeaveID.Eq(uint(leaveID))).Scan(&left); err != nil {
		return err
	}
	// 查詢遞延資訊
	leaveDefer, err := qDefer.WithContext(ctx).Where(qDefer.EmployeeID.Eq(uint(employeeID)), qDefer.LeaveID.Eq(uint(leaveID))).Last()
	if err != nil {
		if !gberror.Is(err, gorm.ErrRecordNotFound) {
			return err
		} else {
			leaveDefer = new(types.LeaveDefer)
		}
	}

	// 計算
	thisTime := left // 這次加的次數
	extra := 0.0     // 超過天數
	left += leaveDefer.Available - leaveDefer.Used
	if left > float64(l.core[employeeID][leaveID].deferDay) {
		extra = left - float64(l.core[employeeID][leaveID].deferDay) // 計算超過天數
		left = float64(l.core[employeeID][leaveID].deferDay)         // 把剩餘天數變為最大天數
	}

	// 插入新遞延結果
	newDefer := &types.LeaveDefer{
		EmployeeID: uint(employeeID),
		LeaveID:    uint(leaveID),
		Effective:  driver.Date(l.now),
		Available:  left,
		Extra:      extra,
		ThisTime:   thisTime,
		PreviousID: leaveDefer.ID,
	}
	if l.core[employeeID][leaveID].isMaturity {
		// 自訂過期?
		newDefer.Expired = driver.Date(time.Now().AddDate(0, int(l.core[employeeID][leaveID].maturity), 0))
	}
	// 將新遞延結果放進資料庫
	if err = qDefer.WithContext(ctx).Create(newDefer); err != nil {
		return err
	}
	// 更新舊遞延
	if _, err = qDefer.WithContext(ctx).Where(qDefer.ID.Eq(leaveDefer.ID)).UpdateSimple(qDefer.NextID.Value(newDefer.ID), qDefer.Expired.Value(driver.Date(l.now))); err != nil {
		return err
	}
	if _, err = qDefer.WithContext(ctx).Where(qDefer.ID.Eq(leaveDefer.ID)).Delete(); err != nil {
		return err
	}

	// commit
	return nil
}

// doUpdate
//
//	@Description: 更新操作
//	@receiver l
func (l *LeaveCheck) doUpdate(ctx context.Context) {
	leaveCorrect := make([]*types.LeaveCorrect, 0)

	// 開啟事務
	err := l.Q.Transaction(func(tx *query.Query) error {
		lc := tx.LeaveCorrect

		for employeeID, c := range l.core {
			for leaveID, info := range c {
				// -----判斷是否到更新日期-----
				if info.update {
					// -------判斷是否遞延-------
					if info.isDefer {
						if err := l.doDefer(ctx, tx, employeeID, leaveID); err != nil {
							return err
						}
					}

					if info.cycle == enum.Calendar && !info.noGroup {
						// 如果是曆年制
						// 判斷年底年資是否滿足下個條件
						if l.endYear[employeeID].year >= int(info.nextCond) {
							info.day = l.calendarCount(l.seniority[employeeID], info.currCond, float64(info.nextDay))
						} else {
							info.day = l.calendarCount(l.seniority[employeeID], info.currCond, info.day)
						}
					}

					leaveCorrect = append(leaveCorrect, &types.LeaveCorrect{
						EmployeeID: uint(employeeID),
						LeaveID:    uint(leaveID),
						Available:  info.day,
					})
					if _, err := lc.WithContext(ctx).Where(lc.EmployeeID.Eq(uint(employeeID)), lc.LeaveID.Eq(uint(leaveID))).Unscoped().Delete(); err != nil {
						return err
					}
				}
			}
		}

		if err := lc.WithContext(ctx).Create(leaveCorrect...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		l.Log.Error(ctx, "save leave correct error", err)
	}
}

// calendarCount
//
//	@Description: 曆年制計算天數
//	@receiver l
//	@param seniority
//	@param now
//	@param nextDay
//	@return float64
func (l *LeaveCheck) calendarCount(seniority seniority, now *types.LeaveGroupCondition, nextDay float64) float64 {
	month := float64(seniority.hireDate.Month() - 1)
	day := float64(seniority.hireDate.Day() - 1)
	maxDay := getDay(seniority.hireDate.Year(), int(seniority.hireDate.Month()))

	// 如果滿足第一個條件
	if seniority.year == 0 {
		// 未滿一年情況
		if seniority.month <= int(now.ConditionNum) {
			return float64(now.Result) + (nextDay - (month+day/maxDay)/12*nextDay)
		}
		// 滿足條件的話要扣掉去年給的天數
		return (month+day/maxDay)/6*float64(now.Result) + (nextDay - (month+day/maxDay)/12*nextDay)
	} else {
		// 滿一年情況
		return (month+day/maxDay)/12*float64(now.Result) + (nextDay - (month+day/maxDay)/12*nextDay)
	}
}

// getDay
//
//	@Description: 獲取當月有幾天
//	@param year
//	@param month
//	@return float64
func getDay(year int, month int) float64 {
	list := []float64{0.0, 31.0, 28.0, 31.0, 30.0, 31.0, 30.0, 31.0, 31.0, 30.0, 31.0, 30.0, 31.0}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		list[2] = 29
	}
	return list[month]
}
