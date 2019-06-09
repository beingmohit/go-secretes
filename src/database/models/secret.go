package models

import (
	"github.com/go-bongo/bongo"
)

type Secret struct {
	bongo.DocumentBase `bson:",inline"`
	Hash string
	SecretText string
	CreatedAt string
	ExpiresAt string
	RemainingViews uint64
}
