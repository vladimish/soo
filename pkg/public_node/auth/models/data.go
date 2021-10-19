package models

import (
	"gorm.io/gorm"
)

type AuthData struct {
	gorm.Model
	NickName      string `json:"nick_name"`
	PublicKey     string `json:"public_key"`
	SecretMessage string `json:"secret_message"`
}
