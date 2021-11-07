package responses

type Responder interface {
	ToJSON() string
}
