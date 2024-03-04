package response

import (
	gbcode "ghostbb.io/gb/errors/gb_code"
	gberror "ghostbb.io/gb/errors/gb_error"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/gin-gonic/gin"
	. "hrapi/internal/utils/response/status"
	"net/http"
)

var RecoveryFn = func(c *gin.Context) {
	if err := recover(); err != nil {
		var gberr error
		switch v := err.(type) {
		case *gberror.Error:
			gberr = v
		default:
			gberr = gberror.NewSkip(1, gbconv.String(v))
		}

		_ = c.Error(gberr)
		Responder(Mount(c)).Fatal(CodeInternalError, gberr.Error())
	}
}

type Responder interface {
	// Ok 狀態碼200，業務成功
	Ok()
	// OkWithDetail 狀態碼200，業務成功，自訂msg輸出
	OkWithDetail(detail string)
	// OkWithData 狀態碼200，業務成功，有 data 實體
	OkWithData(data any)
	// Fail 狀態碼200，業務失敗，標準 error 輸出
	Fail(code gbcode.Code)
	// FailWithMsg 狀態碼200，業務失敗，自定義 msg 輸出
	FailWithMsg(code gbcode.Code, msg string)
	// FailWithDetail 狀態碼200，業務失敗，自定義 error 輸出
	FailWithDetail(code gbcode.Code, detail string)
	// Redirect 狀態碼302，路由重定向
	Redirect(url string)
	// NoRoute 狀態碼404，路由未找到
	NoRoute()
	// Fatal 狀態碼500，服務端致命錯誤，標準 error 輸出 + 自定義詳細信息
	Fatal(code gbcode.Code, detail string)
}

type Ctx struct {
	GinCtx *gin.Context
}

func Mount(c *gin.Context) *Ctx {
	return &Ctx{GinCtx: c}
}

type Response struct {
	Code   int         `json:"code"`
	Data   any         `json:"data"`
	Msg    string      `json:"msg"`
	Detail interface{} `json:"detail"`
}

func result(c *gin.Context, status int, code gbcode.Code, data any) {
	switch status {
	case http.StatusOK:
		c.JSON(status, &Response{
			Code:   code.Code(),
			Msg:    code.Message(),
			Data:   data,
			Detail: code.Detail(),
		})
		break
	case http.StatusNotFound:
		c.AbortWithStatusJSON(status, &Response{
			Code: code.Code(),
			Msg:  code.Message(),
		})
		break
	case http.StatusInternalServerError:
		c.AbortWithStatusJSON(status, &Response{
			Code:   code.Code(),
			Msg:    code.Message(),
			Detail: code.Detail(),
		})
		break
	case http.StatusFound:
		c.Redirect(status, data.(string))
		break
	default:
		c.AbortWithStatusJSON(status, &Response{
			Code:   code.Code(),
			Msg:    code.Message(),
			Detail: code.Detail(),
		})
		break
	}
}

// Ok [200] - 0 - data (is null)
func (c *Ctx) Ok() {
	result(c.GinCtx, http.StatusOK, CodeOK, (*any)(nil))
}

func (c *Ctx) OkWithDetail(detail string) {
	result(c.GinCtx, http.StatusOK, gbcode.New(CodeOK.Code(), CodeOK.Message(), detail), nil)
}

// OkWithData [200] - 0 - data
func (c *Ctx) OkWithData(data any) {
	result(c.GinCtx, http.StatusOK, CodeOK, data)
}

// Fail [200] - 非0 - code
func (c *Ctx) Fail(code gbcode.Code) {
	result(c.GinCtx, http.StatusOK, code, nil)
}

// FailWithMsg c.GinCtx, [200] - 非0 - code&msg (customize)
func (c *Ctx) FailWithMsg(code gbcode.Code, msg string) {
	result(c.GinCtx, http.StatusOK, gbcode.New(code.Code(), msg, code.Detail()), nil)
}

// FailWithDetail [200] - 非0 - code&msg (built-in) detail (customize)
func (c *Ctx) FailWithDetail(code gbcode.Code, detail string) {
	result(c.GinCtx, http.StatusOK, gbcode.New(code.Code(), code.Message(), detail), nil)
}

// Redirect [302] - 非0 - data (redirect)
func (c *Ctx) Redirect(url string) {
	result(c.GinCtx, http.StatusFound, CodeOK, url)
}

// NoRoute [404] - 非0 - code&msg (built-in)
func (c *Ctx) NoRoute() {
	result(c.GinCtx, http.StatusNotFound, CodeNotFound, nil)
}

// Fatal [500] - 非0 - code&msg (built-in) detail (customize)
func (c *Ctx) Fatal(code gbcode.Code, detail string) {
	result(c.GinCtx, http.StatusInternalServerError, gbcode.New(code.Code(), code.Message(), detail), nil)
}
