package customerror

func New(code int, message string) error {
	return &CustomError{code: code,message: message}
}

type CustomError struct {
	code int
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func (e *CustomError) ErrorCode() int {
	return e.code
}