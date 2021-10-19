package network

import (
	"github.com/telf01/soo/pkg/logger"
	"github.com/telf01/soo/pkg/public_node/network/interfaces"
	"github.com/telf01/soo/pkg/public_node/node/models"
	"net/http"
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

func NewContainerSelector() *ContainerSelector{
	cs := &ContainerSelector{}
	cs.Containers = make(map[string]interfaces.RequestContainer)
	return cs
}

func (c *ContainerSelector) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	logger.L.Sugar().Info("Connected")
	for path := range c.Containers {
		if path == r.URL.Path {
			err := c.Containers[path].ParseNext(rw, &r.Body)
			if err != nil{
				logger.L.Sugar().Error(err)
			}
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

func (n *Network) SendMessage(nd *models.Node, msg *interfaces.Responder) error {
	panic("NOT YEY IMPLEMENTED")
}
