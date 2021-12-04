package responses

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/logger"
)

type SearchResult struct {
	URL []string `json:"url" validate:"url"`
}

func (sr *SearchResult) ToJSON() string {
	res, err := json.Marshal(sr)
	if err != nil {
		logger.L.Sugar().Error(err)
		return err.Error()
	}

	return string(res)
}
