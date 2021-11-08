package responses

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/logger"
)

type RegistrationResult struct {
	Result bool `json:"result"`
}

func (rr *RegistrationResult) ToJSON() string {
	res, err := json.Marshal(rr)
	if err != nil {
		logger.L.Sugar().Error(err)
		return err.Error()
	}

	return string(res)
}
