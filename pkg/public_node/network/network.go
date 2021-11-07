package network

import (
	"fmt"
	"github.com/vladimish/soo/pkg/logger"
	"github.com/vladimish/soo/pkg/public_node/network/interfaces"
	"net/http"
	"sync"
)

type Network struct {
	Server *http.Server
	cs     *ContainerSelector
}

func NewNetwork(adr string, s *http.Server) *Network {
	n := Network{
		Server: s,
	}
	n.Server.Addr = adr
	n.cs = NewContainerSelector()
	return &n
}

type ContainerSelector struct {
	Containers map[string]interfaces.RequestContainer
}

func NewContainerSelector() *ContainerSelector {
	cs := &ContainerSelector{}
	cs.Containers = make(map[string]interfaces.RequestContainer)
	return cs
}

func (c *ContainerSelector) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger.L.Sugar().Infof("%s asks on %s", r.Host, r.URL.String())
	for path := range c.Containers {
		if path == r.URL.Path {
			wg := sync.WaitGroup{}
			err := c.Containers[path].ParseNext(rw, &r.Body, &wg)
			if err != nil {
				logger.L.Sugar().Error(err)
			}
			wg.Add(1)
			wg.Wait()
		}
	}
}

func (n *Network) StartListening() error {
	n.Server.Handler = n.cs
	err := n.Server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (n *Network) BindParser(path string, h interfaces.RequestContainer) {
	n.cs.Containers[path] = h
}

func (n *Network) SendMessage(w http.ResponseWriter, msg interfaces.Responder) error {
	fmt.Println(msg.ToJSON())
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(msg.ToJSON()))
	if err != nil {
		return err
	}

	return nil
}
