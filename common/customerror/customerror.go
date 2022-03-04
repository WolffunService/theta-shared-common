package customerror

func New(code int, message string) error {
	return &CustomError{code: code,message: message}
}

func NewCustomError(code int, message string) *CustomError {
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
//ParseToErrorCustom TryParse error to custom error
func ParseToErrorCustom(err error) *CustomError {
	if err == nil{
		return &CustomError{code: 0, message: ""}
	}
	if errCus,ok := err.(*CustomError); ok {
		return errCus
	}
	return &CustomError{code: 0, message: err.Error()}//common.Error == 0
}