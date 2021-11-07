package responses

import "encoding/json"

type DestinationSearch struct {
	RemoteAddress string `json:"remote_address"`
}

func (ds DestinationSearch) ToJSON() string{
	s, err := json.Marshal(ds)
	if err != nil{
		// TODO: Handle error
	}

	return string(s)
}