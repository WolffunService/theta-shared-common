package common

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

func ErrorResponse(code int, message string) *Response {
	debugMessage := ErrorText(code)
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
