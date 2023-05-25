package userrole

import "fmt"

type UserRole uint8

const (
	None UserRole = iota
	Admin
	Root
)

var names = map[UserRole]string{
	None:  "User",
	Admin: "Admin",
	Root:  "Root",
}

func (role UserRole) String() string {
	if name, found := names[role]; found {
		return name
	}

	return fmt.Sprintf("Unknown (%d)", role)
}

func (role UserRole) IsAdmin() bool {
	return role != None
}
