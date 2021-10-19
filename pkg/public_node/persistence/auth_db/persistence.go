package auth_db

import (
	"github.com/telf01/soo/pkg/public_node/auth/models"
)

type Persistence interface {
	GetAuthData(nickName string) (*models.AuthData, error)
	SaveAuth(d *models.AuthData) error
}
