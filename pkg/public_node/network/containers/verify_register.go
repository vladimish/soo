package containers

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/public_node/network/containers/requests"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type VerifyRegisterWrapper struct {
	R  requests.VerifyRegister
	W  http.ResponseWriter
	WG *sync.WaitGroup
}

type VerifyRegister struct {
	waiting chan interface{}
}

func NewVerifyRegister() *VerifyRegister {
	r := VerifyRegister{
		waiting: make(chan interface{}),
	}

	return &r
}

func (r *VerifyRegister) GetChan() chan interface{} {
	return r.waiting
}

// ParseNext creates and adding another VerifyRegisterWrapper object to r.waiting.
func (r *VerifyRegister) ParseNext(writer http.ResponseWriter, reader *io.ReadCloser, wg *sync.WaitGroup) error {
	bytes, err := ioutil.ReadAll(*reader)
	if err != nil {
		return err
	}

	reg := requests.VerifyRegister{}
	err = json.Unmarshal(bytes, &reg)
	if err != nil {
		return err
	}

	rw := VerifyRegisterWrapper{
		R: reg,
		W: writer,
		WG: wg,
	}

	r.waiting <- rw
	return nil
}
