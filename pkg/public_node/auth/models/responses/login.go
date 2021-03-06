package responses

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/logger"
)

type Login struct {
	CheckoutMessage string `json:"checkout_message"`
}

func (l *Login) ToJSON() string {
	res, err := json.Marshal(l)
	if err != nil {
		logger.L.Sugar().Error(err)
		return err.Error()
	}

	return string(res)
}
