package responses

import (
	"encoding/json"
	"github.com/telf01/soo/pkg/logger"
)

type Validation struct {
	EncryptedString string `json:"encrypted_string"`
}

func (v Validation) ToJSON() string {
	s, err := json.Marshal(v)
	if err != nil {
		logger.L.Sugar().Error(err)
	}

	return string(s)
}
