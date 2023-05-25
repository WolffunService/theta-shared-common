package thetancontext

import "fmt"

type ErrUserNotFound struct {
	userID string
	err    any
}

func (e ErrUserNotFound) Error() string {
	return fmt.Sprintf("User %s notfound {{%v}}", e.userID, e.err)
}
