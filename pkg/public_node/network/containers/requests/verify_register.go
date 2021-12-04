package requests

type VerifyRegister struct {
	Nickname        string `json:"nickname" validate:"ascii,gte=4"`
	Signature       string `json:"signature" validate:"base64"`
	CheckoutMessage string `json:"checkout_message" validate:"uuid"`
}
