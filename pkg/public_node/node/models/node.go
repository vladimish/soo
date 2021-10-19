package models

import (
	"image"
)

type Status string

const (
	ACTIVE  Status = "ACTIVE"
	LOST    Status = "LOST"
	OFFLINE Status = "OFFLINE"
)

type Type string

const (
	CLIENT Type = "CLIENT"
	PUBLIC Type = "PUBLIC"
)

type Node struct {
	NickName  string
	PublicKey string
	Hostname  string
	Note      string
	Avatar *image.Image
	Status Status
	Type   Type
}
