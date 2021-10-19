package responses

import "encoding/json"

type Validation struct {
	EncryptedString string `json:"encrypted_string"`
}

func (v Validation) ToJSON() string {
	s, err := json.Marshal(v)
	if err != nil {
		// TODO: Handle error
	}

	return string(s)
}
