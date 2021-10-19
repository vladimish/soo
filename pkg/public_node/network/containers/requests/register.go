package requests

type Register struct {
	NickName string `json:"nick_name"`
	Hostname string `json:"hostname"`
	Key      string `json:"key"`
}
