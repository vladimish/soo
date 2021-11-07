package models

import (
	node_models "github.com/vladimish/soo/pkg/public_node/node/models"
	"gorm.io/gorm"
)

type AuthData struct {
	gorm.Model
	PublicKey       string `json:"public_key"`
	CheckoutMessage string `json:"checkout_message"`
	NodeID          int
	Node            node_models.Node `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
