package model

import (
	"hrapi/internal/types"
	"hrapi/internal/utils/response/page"
)

type GetByKeywordVacationReq struct {
	Keyword string `json:"keyword" from:"keyword"`
	Status  string `json:"status" from:"status"`
}

type GetByKeywordVacationRes struct {
	page.Model[types.Vacation]
}

type GetByIDVacationReq struct {
	ID uint `json:"id"`
}

type GetByIDVacationRes struct {
	*types.Vacation
}

type PostVacationReq struct {
	*types.Vacation
}

type PostVacationRes struct{}

type PutVacationReq struct {
	ID uint `json:"id"`
	*types.Vacation
}

type PutVacationRes struct{}

type DeleteVacationReq struct {
	ID uint `json:"id"`
}

type DeleteVacationRes struct{}

type SetStatusVacationReq struct {
	ID     uint `json:"id"`
	Status bool `json:"status"`
}

type SetStatusVacationRes struct{}

type GetVacationGroupReq struct {
	VacationID uint
}

type GetVacationGroupRes struct {
	*types.VacationGroup
}

type PostVacationGroupReq struct {
	*types.VacationGroup
}

type PostVacationGroupRes struct{}

type DeleteVacationGroupReq struct {
	ID uint
}

type DeleteVacationGroupRes struct{}

type SetVacationGroupNameReq struct {
	ID   uint
	Name string `json:"name"`
}

type SetVacationGroupNameRes struct{}

type SetVacationGroupEmployeeReq struct {
	ID         uint
	EmployeeID []uint `json:"employeeId"`
}

type SetVacationGroupEmployeeRes struct{}

type SetVacationGroupOvertimeRateReq struct {
	ID           uint
	OvertimeRate []*types.VacationGroupOvertimeRate `json:"overtimeRate"`
}

type SetVacationGroupOvertimeRateRes struct{}

type GetByDateVacationScheduleReq struct {
	Start string `form:"start" binding:"required" v:"date-format:Y-m-d"`
	End   string `form:"end" binding:"required" v:"date-format:Y-m-d"`
}

type GetByDateVacationScheduleRes struct {
	*types.VacationSchedule
}

type PostVacationScheduleReq struct {
	*types.VacationScheduleConfig
	Remark     string `json:"remark"`
	VacationID uint   `json:"vacationId"`
}

type PostVacationScheduleRes struct{}

type PutVacationScheduleReq struct {
	GeneralKey string `form:"generalKey" binding:"required"`

	// 內容
	*types.VacationScheduleConfig
	Remark     string `json:"remark"`
	VacationID uint   `json:"vacationId"`
}

type PutVacationScheduleRes struct{}

type DeleteVacationScheduleReq struct {
	ID     uint
	Repeat bool `form:"repeat"`
}

type DeleteVacationScheduleRes struct{}
