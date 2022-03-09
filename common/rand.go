package common

import (
	"crypto/rand"
	"log"
	"math/big"
)

var codeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// RandString random alphabet character with length
func RandString(length int) string {
	var str string
	numCharacters := int64(len(codeAlphabet))
	for i := 0; i < length; i++ {
		str += string(codeAlphabet[cryptoRandSecure(numCharacters)])
	}
	return str
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}
