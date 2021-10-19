package interfaces

import "io"

// RequestContainer handling data parsing and transferring it to a handler with a chan.
type RequestContainer interface {
	GetChan() chan interface{}
	ParseNext(writer io.Writer, reader *io.ReadCloser) error
}
