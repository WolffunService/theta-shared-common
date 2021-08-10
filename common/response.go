package common

import "net/http"

// Response is the response that represents an error.
type Response struct {
	Success       bool        `json:"success,omitempty"`
	Code          int         `json:"code,omitempty"`
	DebugMessage  string      `json:"status,omitempty"`
	RootError     error       `json:"-"`
	RootErrorCode int         `json:"rootCode,omitempty"`
	Message       string      `json:"message,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}

type Paging struct {
	Total int64 `json:"total"`
}

type ResponseWithPaging struct {
	Response `json:",inline"`
	Page     *Paging `json:"page"`
}

// Error root
func (e Response) Root(err error) *Response {
	e.RootError = err
	return &e
}
func (e Response) RootCode(code int) *Response {
	e.RootErrorCode = code
	return &e
}

// Error is required by the error interface.
func (e Response) Error() string {
	return e.Message
}

func ErrorResponse(code int, debugMessage string, message string) *Response {
	return &Response{
		Success:      false,
		Code:         code,
		DebugMessage: debugMessage,
		Message:      message,
	}
}

// SuccessResponse Response Success success and data
func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// SuccessRequest Response Success code and data
func SuccessResponseWithPaging(data interface{}, page *Paging) ResponseWithPaging {
	response := ResponseWithPaging{}
	response.Success = true
	response.Data = data
	response.Page = page
	return response
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg string) Response {
	if msg == "" {
		msg = "We encountered an error while processing your request."
	}
	return Response{
		Code:         http.StatusInternalServerError,
		DebugMessage: http.StatusText(http.StatusInternalServerError),
		Message:      msg,
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg string) Response {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return Response{
		Code:         http.StatusNotFound,
		DebugMessage: http.StatusText(http.StatusNotFound),
		Message:      msg,
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg string) *Response {
	if msg == "" {
		msg = "You are not authenticated to perform the requested action."
	}
	return &Response{
		Code:         http.StatusUnauthorized,
		DebugMessage: http.StatusText(http.StatusUnauthorized),
		Message:      msg,
	}
}
func UnauthorizedRoot(err error) *Response {
	return &Response{
		Code:         http.StatusUnauthorized,
		DebugMessage: http.StatusText(http.StatusUnauthorized),
		Message:      "You are not authenticated to perform the requested action.",
		RootError:    err,
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(err error) *Response {
	return &Response{
		Code:         http.StatusForbidden,
		DebugMessage: http.StatusText(http.StatusForbidden),
		Message:      "You are not authorized to perform the requested action.",
		RootError:    err,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func ErrBadRequest(err error) *Response {
	return &Response{
		Code:         http.StatusBadRequest,
		DebugMessage: http.StatusText(http.StatusBadRequest),
		Message:      "Your request is in a bad format.",
		RootError:    err,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg string) *Response {
	if msg == "" {
		msg = "Your request is in a bad format."
	}
	return &Response{
		Code:         http.StatusBadRequest,
		DebugMessage: http.StatusText(http.StatusBadRequest),
		Message:      msg,
	}
}

func TokenValid(err error, msg string, rootCode int) *Response {
	if msg == "" {
		msg = "Your token is invalid."
	}
	return &Response{
		Code:          http.StatusUnauthorized,
		DebugMessage:  http.StatusText(http.StatusUnauthorized),
		Message:       msg,
		RootErrorCode: rootCode,
		RootError:     err,
	}
}

type invalidField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

