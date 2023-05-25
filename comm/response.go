package comm

import (
	"fmt"

	"github.com/WolffunService/theta-shared-common/thetanerr"
	"github.com/kataras/iris/v12"
)

var Debug = true

type Response struct {
	Success bool                `json:"success"`
	Code    thetanerr.ErrorCode `json:"code"`
	Message string              `json:"message,omitempty"`
	Data    any                 `json:"data,omitempty"`

	Err   error  `json:"-"`
	Debug string `json:"debug,omitempty"`
}

func (r *Response) Preflight(ctx iris.Context) error {
	return nil
}

func (r *Response) Dispatch(ctx iris.Context) {
	err := ctx.JSON(r)
	if err != nil {
		fmt.Println("[Response.Dispatch] error:", err)
	}
}

func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Code:    thetanerr.Success,
		Data:    data,
	}
}

func CodeResponse(code thetanerr.ErrorCode, messages ...string) *Response {
	err := thetanerr.New(code)
	resp := &Response{
		Success: false,
		Code:    code,
	}

	if len(messages) > 0 {
		resp.Message = messages[0]
	} else {
		resp.Message = thetanerr.ErrMessage(err)
	}

	return resp
}

func CodeResponseWithParams(code thetanerr.ErrorCode, params ...any) *Response {
	err := thetanerr.New(code)
	resp := &Response{
		Success: false,
		Code:    code,
	}

	patternMessage := thetanerr.ErrMessage(err)
	if params != nil && len(params) >= 1 {
		patternMessage = fmt.Sprintf(patternMessage, params...)
	}

	resp.Message = patternMessage

	return resp
}

func ErrorResponse(err error) *Response {
	code := thetanerr.ErrCode(err)
	message := thetanerr.ErrMessage(err)

	resp := &Response{
		Success: false,
		Code:    code,
		Err:     err,
		Message: message,
	}

	// Build debug error message
	if Debug {
		resp.Debug = fmt.Sprintf("Err: %+v", err)
	}

	return resp
}

func ErrorResponseCustomMsg(err error, message string) *Response {
	code := thetanerr.ErrCode(err)

	resp := &Response{
		Success: false,
		Code:    code,
		Err:     err,
		Message: message,
	}

	// Build debug error message
	if Debug {
		resp.Debug = fmt.Sprintf("Err: %+v", err)
	}

	return resp
}

// AutoResponse Automatically response by handling error
// and return adapt response correctly
func AutoResponse(data any, err error) *Response {
	if err != nil {
		return ErrorResponse(err)
	}

	return SuccessResponse(data)
}
