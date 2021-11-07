package requests

type VerifyRegister struct {
	Nickname        string `json:"nickname"`
	Signature       string `json:"signature"`
	CheckoutMessage string `json:"checkout_message"`
}
