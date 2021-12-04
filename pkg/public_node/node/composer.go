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
	h[REGISTER] = containers.NewRegister()
	h[VERIFY_REGISTER] = containers.NewVerifyRegister()
	h[FIND_USER] = containers.NewFindUser()
	c.cm.BindConnectionHandlers(h)

	go c.cm.HandleRegister(h[REGISTER].GetChan())
	go c.cm.HandleVerifyRegister(h[VERIFY_REGISTER].GetChan())
	go c.cm.HandleFindUser(h[FIND_USER].GetChan())

	c.cm.StartServer()
}
