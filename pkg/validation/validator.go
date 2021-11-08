package validation

import (
	"encoding/base64"
	"github.com/go-playground/validator/v10"
)

var V = validator.New()

func init() {
	V.RegisterValidation("ed25519_public_key", func(fl validator.FieldLevel) bool {
		base64Key := fl.Field().String()
		res, err := base64.StdEncoding.DecodeString(base64Key)
		if len(res) != 32 || err != nil {
			return false
		}

		return true
	})
}
