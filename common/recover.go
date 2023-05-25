package common

import (
	"log"

	"github.com/go-errors/errors"
)

func Recover(message ...string) {
	if err := recover(); err != nil {
		log.Println("recovered mes: ", message, " -- from: ", errors.Wrap(err, 2).ErrorStack())
	}
}
