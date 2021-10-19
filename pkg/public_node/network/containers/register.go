package containers

import (
	"encoding/json"
	"github.com/telf01/soo/pkg/public_node/network/containers/requests"
	"io"
	"io/ioutil"
)

type RegisterWrapper struct {
	R requests.Register
	W io.Writer
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
func (r *Register) ParseNext(writer io.Writer, reader *io.ReadCloser) error {
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
	}

	r.waiting <- rw
	return nil
}
