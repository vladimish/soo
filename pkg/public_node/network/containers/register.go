package containers

import (
	"encoding/json"
	"github.com/telf01/soo/pkg/public_node/network/containers/requests"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type RegisterWrapper struct {
	R  requests.Register
	W  http.ResponseWriter
	WG *sync.WaitGroup
}

type Register struct {
	waiting chan interface{}
}

func NewRegister() *Register {
	r := Register{
		waiting: make(chan interface{}),
	}

	return &r
}

func (r *Register) GetChan() chan interface{} {
	return r.waiting
}

// ParseNext creates and adding another RegisterWrapper object to r.waiting.
func (r *Register) ParseNext(writer http.ResponseWriter, reader *io.ReadCloser, wg *sync.WaitGroup) error {
	bytes, err := ioutil.ReadAll(*reader)
	if err != nil {
		return err
	}

	reg := requests.Register{}
	err = json.Unmarshal(bytes, &reg)
	if err != nil {
		return err
	}

	rw := RegisterWrapper{
		R: reg,
		W: writer,
		WG: wg,
	}

	r.waiting <- rw
	return nil
}
