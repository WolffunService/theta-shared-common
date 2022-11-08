package common

import (
	"github.com/go-errors/errors"
	"log"
)

type UserRole int

const (
	NONE     UserRole = 0 // Role Default
	ADMIN    UserRole = 1 // Role Admin
	ROOT     UserRole = 2 // Role Highest
	SYSADMIN UserRole = 3 // Role System Config
)

func Recover(message ...string) {
	if err := recover(); err != nil {
		log.Println("recovered mes: ", message, " -- from: ", errors.Wrap(err, 2).ErrorStack())
	}
}
