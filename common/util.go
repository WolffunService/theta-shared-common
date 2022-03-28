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

func Recover() {
	if err := recover(); err != nil {
		log.Println("recovered from ", errors.Wrap(err, 2).ErrorStack())
	}
}
