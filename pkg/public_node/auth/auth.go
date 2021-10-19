package auth

import (
	"github.com/telf01/soo/pkg/public_node/network/interfaces"
	"github.com/telf01/soo/pkg/public_node/node/models"
	"github.com/telf01/soo/pkg/public_node/persistence/auth_db"
)

type Auth struct {
	p auth_db.Persistence
}

func NewAuth(p auth_db.Persistence) *Auth {
	a := Auth{
		p: p,
	}

	return &a
}

func (a *Auth) GetNodeOrNil(nn string) (*models.Node, error){
	// TODO: Add real db check.
	return nil, nil
}

func (a *Auth) CheckAuth(n *models.Node) {
	panic("NOT YET IMPLEMENTED")
}

// BuildAuthMessage creates new response based on current user status.
func (a *Auth) BuildAuthMessage(node *models.Node) (*interfaces.Responder, error) {
	ad, err := a.p.GetAuthData(node.NickName)
	if err != nil{
		return nil, err
	}
	if ad == nil{

	}

	return nil, nil
}

func CreateAuth(n *models.Node) {
	panic("NOT YET IMPLEMENTED")
}
