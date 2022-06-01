package common

import (
	"github.com/go-errors/errors"
	"log"
)

type UserRole int

const (
	NONE  UserRole = 0
	ADMIN UserRole = 1
	ROOT  UserRole = 2
)

func Recover(message ...string) {
	if err := recover(); err != nil {
		log.Println("recovered mes: ", message, " -- from: ", errors.Wrap(err, 2).ErrorStack())
	}
}
