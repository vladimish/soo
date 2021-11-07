package node

import (
	"github.com/vladimish/soo/pkg/public_node/network/containers"
	"github.com/vladimish/soo/pkg/public_node/network/interfaces"
)

// Composer is a high level object meant to maintain high level configuration.
type Composer struct {
	cm *ConnectionManager
}

func NewComposer(cm *ConnectionManager) *Composer {
	c := Composer{
		cm: cm,
	}

	return &c
}

func (c *Composer) Start() {
	var h = make(map[string]interfaces.RequestContainer)
	h[Register] = containers.NewRegister()
	h[VerifyRegister] = containers.NewVerifyRegister()
	c.cm.BindConnectionHandlers(h)

	go c.cm.HandleRegister(h[Register].GetChan())
	go c.cm.HandleVerifyRegister(h[VerifyRegister].GetChan())

	c.cm.StartServer()
}
