package customerror

import "github.com/WolffunGame/theta-shared-common/common"

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
//ParseToErrorCustom TryParse error to custom error
func ParseToErrorCustom(err error) *CustomError {
	if errCus,ok := err.(*CustomError); ok {
		return errCus
	}
	return &CustomError{code: common.Error, message: err.Error()}
}