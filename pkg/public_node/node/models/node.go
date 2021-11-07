package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	REGISTRATION Status = "REGISTRATION"
	ACTIVE       Status = "ACTIVE"
	LOST         Status = "LOST"
	OFFLINE      Status = "OFFLINE"
)

type Type string

const (
	CLIENT Type = "CLIENT"
	PUBLIC Type = "PUBLIC"
)

type Node struct {
	gorm.Model
	NickName   string `gorm:"size:16; unique"`
	Note       string
	AvatarPath uuid.UUID
	Status     Status
	Type       Type
}
