package api

import (
	gbhttp "ghostbb.io/gb/net/gb_http"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/salary/v1/model"
	"hrapi/apps/hr/salary/v1/service"
	"hrapi/internal/middleware"
	"hrapi/internal/utils/paginator"

	. "hrapi/internal/utils/response"
	. "hrapi/internal/utils/response/status"
)

type SalaryApi struct{}

func (s *SalaryApi) Init(group *gin.RouterGroup) {
	// 需要有後台權限 (middleware.Software())
	v1 := group.Group("salary").Use(middleware.Auth(), middleware.Software())
	v1.POST("work", s.calc)
	v1.GET("", middleware.Paginator(), s.getSalaryCalc)
	v1.DELETE(":id", s.delete)
	v1.GET(":id", s.getByID)
	v1.PUT("calcEmployee", s.saveCalcEmployee)
}

// 薪資作業
//
//	route => POST /api/v1/salary/work
func (s *SalaryApi) calc(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SalaryCalcReq
		err error
	)

	in.FounderEmployeeID = gbhttp.Get(c, "employee.id").Uint()

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Salary(ctx).Calc(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 查詢所有薪資作業
//
//	route => GET /api/v1/salary
func (s *SalaryApi) getSalaryCalc(c *gin.Context) {
	var (
		ctx  = gbhttp.Ctx(c)
		out  model.GetSalaryCalcRes
		page = gbhttp.Get(c, "paginator").Interface().(*paginator.Pagination)
		err  error
	)

	out, err = service.Salary(ctx, page).GetSalaryCalc()
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 刪除薪資作業
//
//	route => DELETE /api/v1/salary/:id
func (s *SalaryApi) delete(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.DeleteSalaryCalcReq
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	err = service.Salary(ctx).Delete(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}

// 根據ID查詢薪資作業
//
//	route => GET /api/v1/salary/:id
func (s *SalaryApi) getByID(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.GetByIDSalaryCalcReq
		out model.GetByIDSalaryCalcRes
		err error
	)

	in.ID = gbconv.Uint(c.Param("id"))

	out, err = service.Salary(ctx).GetByIDSalaryCalc(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).OkWithData(out)
}

// 儲存設定
//
//	route => PUT /api/v1/salary/calcEmployee
func (s *SalaryApi) saveCalcEmployee(c *gin.Context) {
	var (
		ctx = gbhttp.Ctx(c)
		in  model.SaveEmployeeItemReq
		err error
	)

	if err = gbhttp.ParseJSON(c, &in); err != nil {
		Responder(Mount(c)).FailWithDetail(CodeRequestInvalidBody, err.Error())
		return
	}

	err = service.Salary(ctx).SaveCalcEmployee(in)
	if err != nil {
		Responder(Mount(c)).FailWithMsg(CodeFailed, err.Error())
		return
	}

	Responder(Mount(c)).Ok()
}
