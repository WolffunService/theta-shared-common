package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"github.com/WolffunGame/theta-shared-common/common"
	"github.com/WolffunGame/theta-shared-common/common/thetaerror"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/rs/zerolog/log"
)

func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	pubkey_pem := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(pubkey)}))
	return pubkey_pem
}
func ExportPrivateKeyAsPemStr(privatekey *rsa.PrivateKey) string {
	privatekey_pem := string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}))
	return privatekey_pem
}
func ExportMsgAsPemStr(msg []byte) string {
	msg_pem := string(pem.EncodeToMemory(&pem.Block{Type: "MESSAGE", Bytes: msg}))
	return msg_pem
}
func main() {
	bits := 1024
	flag.Parse()
	args := flag.Args()

	m := args[0]

	bobPrivateKey, _ := rsa.GenerateKey(rand.Reader, bits)

	bobPublicKey := &bobPrivateKey.PublicKey

	fmt.Printf("%s\n", ExportPrivateKeyAsPemStr(bobPrivateKey))

	fmt.Printf("%s\n", ExportPublicKeyAsPemStr(bobPublicKey))

	message := []byte(m)
	label := []byte("123123")
	hash := sha256.New()

	ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, bobPublicKey, message, label)

	fmt.Printf("%s\n", ExportMsgAsPemStr(ciphertext))

	plainText, _ := rsa.DecryptOAEP(hash, rand.Reader, bobPrivateKey, ciphertext, label)

	fmt.Printf("RSA decrypted to [%s]", plainText)

	////////////How to use json logger
	debug := false
	// Apply log level in the beginning of the application
	thetalog.SetGlobalLevel(thetalog.InfoLevel)
	if debug {
		thetalog.SetGlobalLevel(thetalog.DebugLevel)
	}

	thetalog.Info().
		Str("service", "my-service").
		Int("Some integer", 10).
		Msg("Hello")
	// Debug log
	log.Debug().Msg("Exiting Program")

	////////////How to use json logger with default value
	logger := thetalog.With().Str("service", "theta-data").
		Str("node", "localhost").
		Logger()

	logger.Err(&thetaerror.Error{
		Code:    common.BusyServer,
		Message: "Server is too busy",
		Op:      "Convert",
		Err:     nil,
	})
}
