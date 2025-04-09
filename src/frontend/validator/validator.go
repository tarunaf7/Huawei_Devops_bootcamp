













package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate




func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Payload interface {
	Validate() error
}

type AddToCartPayload struct {
	Quantity  uint64 `validate:"required,gte=1,lte=10"`
	ProductID string `validate:"required"`
}

type PlaceOrderPayload struct {
	Email         string `validate:"required,email"`
	StreetAddress string `validate:"required,max=512"`
	ZipCode       int64  `validate:"required"`
	City          string `validate:"required,max=128"`
	State         string `validate:"required,max=128"`
	Country       string `validate:"required,max=128"`
	CcNumber      string `validate:"required,credit_card"`
	CcMonth       int64  `validate:"required,gte=1,lte=12"`
	CcYear        int64  `validate:"required"`
	CcCVV         int64  `validate:"required"`
}

type SetCurrencyPayload struct {
	Currency string `validate:"required,iso4217"`
}


func (ad *AddToCartPayload) Validate() error {
	return validate.Struct(ad)
}

func (po *PlaceOrderPayload) Validate() error {
	return validate.Struct(po)
}

func (sc *SetCurrencyPayload) Validate() error {
	return validate.Struct(sc)
}


func ValidationErrorResponse(err error) error {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors.New("invalid validation error format")
	}
	var msg string
	for _, err := range validationErrs {
		msg += fmt.Sprintf("Field '%s' is invalid: %s\n", err.Field(), err.Tag())
	}
	return fmt.Errorf(msg)
}
