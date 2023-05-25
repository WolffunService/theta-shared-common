package iapvalidator

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/WolffunService/theta-shared-common/enums/platformtypeenum"
	"github.com/WolffunService/theta-shared-common/iapvalidator/iapenum"
	"github.com/awa/go-iap/appstore"
	"github.com/awa/go-iap/playstore"
	"github.com/buger/jsonparser"
)

var (
	ErrPlatformTypeInvalid     = errors.New("platformType invalid")
	ErrPackIdNotFound          = errors.New("packId not found")
	ErrReceiptDataInvalid      = errors.New("receipt data invalid")
	ErrReceiptDataAlreadyExist = errors.New("receipt data already exist")
)

type ValidatorController interface {
	HandleErr(userId, reason string) error                                                            //handle error - reciept data invalid
	GetPackByProductId(productId string) (IPackItem, error)                                           //get pack ID from IAPCode(ProductID)
	CheckExitsReceipt(platformType platformtypeenum.PlatformType, transactionId string) (bool, error) //validate storage receipt
	CreateNewReceipt(ctx context.Context, userId string, buyRequest BuyItemReq) error                 //handle create new receipt data after flow success
}

type IPackItem interface {
	GetPackID() iapenum.PackID
	GetPackType() int
	GetPriceUSD() float64
	GetItems() any //items received
	IsRealMoneyPack() bool
	GetIAPCode() string
}

func IsSamePack(packA, packB IPackItem) bool {
	return packA.GetIAPCode() == packB.GetIAPCode() &&
		packA.GetPackType() == packB.GetPackType() &&
		packA.GetPackID() == packB.GetPackID()
}

// PurchaseState is the data type for purchase states.
type PurchaseState int

// List of purchase states.
const (
	PurchaseDone     PurchaseState = 0
	PurchaseCanceled PurchaseState = 1
	PurchasePending  PurchaseState = 2
)

type BuyItemReq struct {
	PackId       int                           `json:"packId" form:"packId"`
	PlatformType platformtypeenum.PlatformType `json:"platformType" validate:"enum"`
	ReceiptJson  string                        `json:"receiptJson"`
	Signature    string                        `json:"signature"`

	bytesReceipt []byte
}

func (r BuyItemReq) GetTransactionId() (string, error) {
	switch r.PlatformType {
	case platformtypeenum.ANDROID:
		return r.GetString("orderId")
	case platformtypeenum.IOS:
		return r.GetString("TransactionID")
	}
	return "", errors.New("can't get transactionId from receipt")
}

func (r BuyItemReq) getBytesReceipt() []byte {
	if len(r.bytesReceipt) == 0 {
		r.bytesReceipt = []byte(r.ReceiptJson)
	}
	return r.bytesReceipt
}

func (r BuyItemReq) GetString(keys ...string) (string, error) {
	return jsonparser.GetString(r.getBytesReceipt(), keys...)
}

type Validator struct {
	CompanyName            string
	ProductName            string
	Base64EncodedPublicKey string // to verify signature in GooglePlay
	//base64JsonKey          string         // to init jsonKey
	//jsonKey                []byte         // to test
	ValidatorController
}

func NewValidator(companyName, productName, publicKey string, controller ValidatorController) Validator {
	if companyName == "" || productName == "" || publicKey == "" || controller == nil {
		panic("Validator invalid")
	}

	return Validator{
		CompanyName:            companyName,
		ProductName:            productName,
		Base64EncodedPublicKey: publicKey,
		ValidatorController:    controller,
	}
}

func (v Validator) VerifySignatureGooglePlay(receipt []byte, signature string) (bool, IPackItem, error) {
	isValid, err := playstore.VerifySignature(v.Base64EncodedPublicKey, receipt, signature)

	if !isValid || err != nil {
		return false, nil, err
	}

	productId, errParseJson := jsonparser.GetString(receipt, "productId")
	if errParseJson != nil {
		return false, nil, fmt.Errorf("productId invalid : %v %v ", errParseJson, productId)
	}

	purchaseState, errParseJson := jsonparser.GetInt(receipt, "purchaseState")
	if errParseJson != nil || PurchaseState(purchaseState) != PurchaseDone {
		return false, nil, fmt.Errorf("purchaseState invalid : %v %v ", errParseJson, purchaseState)
	}

	pack, errGetPackId := v.GetPackByProductId(productId)
	if errGetPackId != nil {
		return false, nil, errGetPackId
	}

	return isValid, pack, err
}

// VerifyAppstore status == 0 is valid
// https://developer.apple.com/documentation/appstorereceipts/status
func (v Validator) VerifyAppstore(ctx context.Context, receipt []byte) (*appstore.IAPResponse, bool, error) {
	payload, err := jsonparser.GetString(receipt, "Payload")
	if err != nil {
		return nil, false, err
	}

	client := appstore.New()
	request := appstore.IAPRequest{
		ReceiptData: payload,
	}
	response := &appstore.IAPResponse{}

	err = client.Verify(ctx, request, response)
	return response, response.Status == 0, err
}

func (v Validator) VerifyIAP(ctx context.Context, userId string, req BuyItemReq) (bool, IPackItem, error) {
	platformType := req.PlatformType
	if !platformType.IsValid() {
		return false, nil, ErrPlatformTypeInvalid
	}

	transId, errGetTransId := req.GetTransactionId()
	if errGetTransId != nil {
		return false, nil, errGetTransId
	}

	receiptExits, err := v.CheckExitsReceipt(req.PlatformType, transId)
	if err != nil {
		return false, nil, err
	}

	if receiptExits {
		//log.Printf("[VerifyIAP] [WARNING] Detect userId: %s bug mutil claim receipt in game store.", userId)
		return false, nil, ErrReceiptDataAlreadyExist
	}

	switch platformType {
	case platformtypeenum.ANDROID:
		isValid, pack, errPlayStore := v.VerifySignatureGooglePlay(req.getBytesReceipt(), req.Signature)
		if errPlayStore == ErrPackIdNotFound {
			return false, nil, v.err(userId, req.ReceiptJson)
		}
		return isValid, pack, errPlayStore
	case platformtypeenum.IOS:
		res, statusOk, errV := v.VerifyAppstore(ctx, req.getBytesReceipt())

		//duple check -- load packId
		if errV == nil && statusOk {
			if len(res.Receipt.InApp) == 0 {
				return false, nil, v.err(userId, req.ReceiptJson)
			}
			split := strings.Split(res.Receipt.InApp[0].ProductID, ".")
			if len(split) <= 2 {
				return false, nil, v.err(userId, req.ReceiptJson)
			}
			if !(split[1] == v.CompanyName && split[2] == v.ProductName) {
				return false, nil, v.err(userId, req.ReceiptJson)
			}

			productId := res.Receipt.InApp[0].ProductID
			pack, errGetPackId := v.GetPackByProductId(productId)

			if errGetPackId != nil {
				return false, nil, v.err(userId, req.ReceiptJson)
			}
			return statusOk, pack, errV
		}
		return statusOk, nil, errV
	}
	return false, nil, ErrPlatformTypeInvalid
}

func (v Validator) err(userId, reason string) error {
	if err := v.HandleErr(userId, reason); err != nil {
		return err
	}
	return ErrReceiptDataInvalid
}
