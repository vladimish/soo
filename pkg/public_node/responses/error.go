package responses

import (
	"encoding/json"
	"github.com/telf01/soo/pkg/logger"
)

type ErrorCode string

const (
	BAD_REQUEST ErrorCode = "BAD_REQUEST"
	TIMEOUT     ErrorCode = "TIMEOUT"
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e Error) ToJSON() string {
	s, err := json.Marshal(e)
	if err != nil {
		logger.L.Sugar().Error(err)
	}

	return string(s)
}
