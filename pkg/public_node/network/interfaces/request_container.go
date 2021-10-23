package interfaces

import (
	"io"
	"net/http"
	"sync"
)

// RequestContainer handling data parsing and transferring it to a handler with a chan.
type RequestContainer interface {
	GetChan() chan interface{}
	ParseNext(writer http.ResponseWriter, reader *io.ReadCloser, wg *sync.WaitGroup) error
}
