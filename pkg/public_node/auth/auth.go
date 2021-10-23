package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	models2 "github.com/telf01/soo/pkg/public_node/auth/models"
	"github.com/telf01/soo/pkg/public_node/auth/models/responses"
	"github.com/telf01/soo/pkg/public_node/network/interfaces"
	"github.com/telf01/soo/pkg/public_node/node/models"
	"github.com/telf01/soo/pkg/public_node/persistence/auth_db"
	"gorm.io/gorm"
	"io"
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

func (a *Auth) CheckAuth(n *models.Node) {
	panic("NOT YET IMPLEMENTED")
}

// BuildAuthMessage builds new message, based on AuthData.
func (a *Auth) BuildAuthMessage(ad models2.AuthData) (interfaces.Responder, error) {
	block, err := aes.NewCipher([]byte(calculateMD5(ad.SecretMessage)))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(ad.SecretMessage), nil)
	r := &responses.Login{
		EncryptedMessage: string(ciphertext),
	}
	return r, nil
}

func calculateMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (a *Auth) CreateAuth(node *models.Node, pk string) (*models2.AuthData, error) {
	sm, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	ta := &models2.AuthData{
		PublicKey:     pk,
		SecretMessage: sm.String(),
		Node:          *node,
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
