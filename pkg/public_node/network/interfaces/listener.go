package interfaces

type Listener interface {
	StartListening() error
	BindHandler(path string, h *RequestContainer)
}
