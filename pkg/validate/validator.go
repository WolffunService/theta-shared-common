package validate

import (
	"sync"

	"github.com/WolffunService/thetan-shared-common/enums"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Ins   *validator.Validate
	Uni   *ut.UniversalTranslator
	Trans ut.Translator
)

var loadOnce sync.Once

func Init() {
	loadOnce.Do(func() {
		enTrans := en.New()
		Uni = ut.New(enTrans)
		Trans, _ = Uni.GetTranslator(enTrans.Locale())

		Ins = validator.New()
		err := en_translations.RegisterDefaultTranslations(Ins, Trans)
		if err != nil {
			panic(err)
		}

		err = Ins.RegisterValidation("object-id", ValidateObjectId)
		if err != nil {
			panic(err)
		}

		err = Ins.RegisterValidation("enum", enums.ValidateEnum, true)
		if err != nil {
			panic(err)
		}
	})
}

func ValidateObjectId(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return primitive.IsValidObjectID(str)
}
