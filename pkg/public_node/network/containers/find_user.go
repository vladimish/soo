package containers

import (
	"encoding/json"
	"github.com/vladimish/soo/pkg/public_node/network/containers/requests"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type FindUserWrapper struct {
	R  requests.FindUser
	W  http.ResponseWriter
	WG *sync.WaitGroup
}

type FindUser struct {
	waiting chan interface{}
}

func NewFindUser() *FindUser {
	f := FindUser{
		waiting: make(chan interface{}),
	}

	return &f
}

func (f *FindUser) GetChan() chan interface{} {
	return f.waiting
}

// ParseNext creates and adding another FindUserWrapper object to r.waiting.
func (f *FindUser) ParseNext(writer http.ResponseWriter, reader *io.ReadCloser, wg *sync.WaitGroup) error {
	bytes, err := ioutil.ReadAll(*reader)
	if err != nil {
		return err
	}

	fnd := requests.FindUser{}
	err = json.Unmarshal(bytes, &fnd)
	if err != nil {
		return err
	}

	rw := FindUserWrapper{
		R:  fnd,
		W:  writer,
		WG: wg,
	}

	f.waiting <- rw
	return nil
}
