package interfaces

// Responder is a base for formatted server response.
type Responder interface {
	ToJSON() string
}