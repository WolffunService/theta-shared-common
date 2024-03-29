package common

import (
	"log"

	"github.com/go-errors/errors"
)

type UserRole int

const (
	NONE     UserRole = 0 // Role Default
	ADMIN    UserRole = 1 // Role Admin
	ROOT     UserRole = 2 // Role Highest
	SYSADMIN UserRole = 3 // Role System Config
	EDITOR   UserRole = 4 // Role Editor
)

func Recover(message ...string) {
	if err := recover(); err != nil {
		log.Println("recovered mes: ", message, " -- from: ", errors.Wrap(err, 2).ErrorStack())
	}
}
