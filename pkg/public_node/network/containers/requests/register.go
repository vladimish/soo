package requests

type Register struct {
	Nickname string `json:"nickname" validate:"ascii,gte=4"`
	Key      string `json:"key" validate:"base64,ed25519_public_key"`
}
