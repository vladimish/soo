package responses

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/logger"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SearchResult struct {
	Results []Result `json:"results"`
}

func (sr *SearchResult) ToJSON() string {
	res, err := json.Marshal(sr)
	if err != nil {
		logger.L.Sugar().Error(err)
		return err.Error()
	}

	return string(res)
}
