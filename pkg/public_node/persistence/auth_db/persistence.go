package auth_db

import (
	"github.com/vladimish/soo/pkg/public_node/auth/models"
	node_models "github.com/vladimish/soo/pkg/public_node/node/models"
)

type Persistence interface {
	GetNode(nickName string) (*node_models.Node, error)
	SaveNode(node *node_models.Node) error
	GetAuthData(message string) (*models.AuthData, error)
	GetLastAuthData(node node_models.Node) (*models.AuthData, error)
	SaveAuth(d *models.AuthData) error
}
