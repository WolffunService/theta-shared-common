package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

//0x98188a7526c5eb0202a8cc5ac16861ad412e6072
func TestAWT(t *testing.T) {
	x, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		ClaimKeySid:     "61a9b9dd003b2a9503350a91",
		ClaimKeySub:     "asd",
		ClaimKeyId:      "61a9b9dd003b2a9503350a91",
		ClaimKeyRole:    2,
		ClaimKeyCanMint: false,
		ClaimKeyNbf:     time.Now().Unix(),
		ClaimKeyIss:     "https://api.marketplace.app",
		ClaimKeyAud:     "JWT_APIS",
		ClaimKeyExp:     time.Now().Add(time.Second * time.Duration(604800)).Unix(),
	}).SignedString([]byte("_Token---1234567890_APIS_SECRET"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(x)
}
