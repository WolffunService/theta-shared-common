package common
const (

	//common code 1 - 1000
	TokenInvalid = 1
	TokenExpired = 2

	//wolffunId ERROR_CODE 2001 -> 3000:
	EmailNotExist = 2001
	EmailIsExist  = 2002
	CodeNotValid  = 2003
	EmailNotValid = 2004

)
var errorText = map[int]string{
	//common code 1 - 1000
	TokenInvalid: "Token Invalid",
	TokenExpired: "Token Expired",

	//wolffunId ERROR_CODE 2001 -> 3000:
	EmailNotExist: "Email Not Exist In Database",
	EmailIsExist: "Email Exist In Database",
	CodeNotValid: "Code Not Valid",
	EmailNotValid: "Email Not Valid",
}
// StatusText returns a text for the common error code. It returns the empty
// string if the code is unknown.
func ErrorText(code int) string {
	return errorText[code]
}