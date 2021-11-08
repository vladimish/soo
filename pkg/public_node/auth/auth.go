package auth

import (
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	models2 "github.com/vladimish/soo/pkg/public_node/auth/models"
	"github.com/vladimish/soo/pkg/public_node/auth/models/responses"
	"github.com/vladimish/soo/pkg/public_node/network/interfaces"
	"github.com/vladimish/soo/pkg/public_node/node/models"
	"github.com/vladimish/soo/pkg/public_node/persistence/auth_db"
	"golang.org/x/crypto/ed25519"
	"gorm.io/gorm"
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

func (a *Auth) GetNodeOrNil(nn string) (*models.Node, error) {
	n, err := a.p.GetNode(nn)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return n, nil
}

func (a *Auth) CheckAuth(signature string, message string) (bool, error) {
	ad, err := a.p.GetAuthData(message)
	if err != nil {
		return false, err
	}

	// Unwrap base64 signature to byte array
	sigBytes := make([]byte, 64)
	base64.StdEncoding.Decode(sigBytes, []byte(signature))
	pkBytes := make([]byte, 32)
	base64.StdEncoding.Decode(pkBytes, []byte(ad.PublicKey))

	return ed25519.Verify(pkBytes, []byte(message), sigBytes), nil
}

// BuildAuthMessage builds new message, based on AuthData.
func (a *Auth) BuildAuthMessage(ad models2.AuthData) (interfaces.Responder, error) {
	r := &responses.Login{
		CheckoutMessage: ad.CheckoutMessage,
	}
	return r, nil
}

func (a *Auth) CreateAuth(node *models.Node, pk string) (*models2.AuthData, error) {
	sm, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	ta := &models2.AuthData{
		PublicKey:       pk,
		CheckoutMessage: sm.String(),
		Node:            *node,
	}
	err = a.p.SaveAuth(ta)
	if err != nil {
		return nil, err
	}
	return ta, nil
}

func (a *Auth) SaveNode(node *models.Node) error {
	return a.p.SaveNode(node)
}

func (a *Auth) ChangeNodeStatus(message string, status models.Status) error {
	ad, err := a.p.GetAuthData(message)
	if err != nil {
		return err
	}

	err = a.p.UpdateNode(ad.NodeID, "status", models.ACTIVE)
	if err != nil {
		return err
	}

	return nil
}
