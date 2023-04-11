package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/WolffunService/theta-shared-common/thetalog"
	"time"
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
	////////////How to use json logger
	debug := false
	// Apply log level in the beginning of the application
	thetalog.SetGlobalLevel(thetalog.InfoLevel)
	if debug {
		thetalog.SetGlobalLevel(thetalog.DebugLevel)
	}

	//thetalog.Info().
	//	Str("service", "my-service").
	//	Int("Some integer", 10).
	//	Msg("Hello")
	//// Debug log
	//log.Debug().Msg("Exiting Program")
	//
	//////////////How to use json logger with default value
	//logger := thetalog.With().Str("service", "theta-data").
	//	Str("node", "localhost").
	//	Logger()
	//
	//logger.Err(&thetaerror.Error{
	//	Code:    common.BusyServer,
	//	Message: "Server is too busy",
	//	Op:      "Convert",
	//	Err:     nil,
	//})
	//
	////How to use eventbus
	//handler := func(a, b int) {
	//	time.Sleep(3 * time.Second)
	//	fmt.Printf("Event handler")
	//	fmt.Printf("%d\n", a+b)
	//}
	//
	//bus := EventBus.New()
	//err := bus.SubscribeAsync("main:slow_calculator", handler, false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//bus.Publish("main:slow_calculator", 20, 60)
	//time.Sleep(5 * time.Second)
	//bus.Publish("main:slow_calculator", 10, 20)
	//
	//fmt.Println("start: do some stuff while waiting for a result")
	//fmt.Println("end: do some stuff while waiting for a result")
	//
	////bus.WaitAsync() // wait for all async callbacks to complete
	//
	//time.Sleep(20 * time.Second)
	//
	//fmt.Println("do some stuff after waiting for result")

	defer TimeTrack(time.Now(), "main")
	logger := thetalog.NewBizLogger("bizname")
	logger.Log().Op("bussinessName")
	for i := 0; i < 100000; i++ {

		logger.Err(fmt.Errorf("error ")).Op("main").Int("z", 1).Msg("hahahaaaaa")
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", name, elapsed)
}
