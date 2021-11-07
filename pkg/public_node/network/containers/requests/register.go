package requests

type Register struct {
	Nickname string `json:"nickname"`
	Key      string `json:"key"`
}
