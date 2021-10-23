package responses

import (
	"encoding/json"
	"github.com/telf01/soo/pkg/logger"
	"strconv"
	"strings"
)

type Login struct {
	EncryptedMessage string `json:"encrypted_message"`
}

func (l *Login) ToJSON() string {
	res, err := json.Marshal(l)
	if err != nil {
		logger.L.Sugar().Error(err)
		return err.Error()
	}
	rawRes, err := unescapeUnicodeCharactersInJSON(res)
	return string(rawRes)
}

// TODO: Move to utils.
func unescapeUnicodeCharactersInJSON(jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}