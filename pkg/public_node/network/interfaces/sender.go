package interfaces

import (
	"github.com/telf01/soo/pkg/public_node/node/models"
)

type Sender interface {
	SendMessage(n *models.Node, msg *Responder) error
}
