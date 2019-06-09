package models

import (
	"github.com/go-bongo/bongo"
)

type Secret struct {
	bongo.DocumentBase `bson:",inline"`
	Hash string `json:"hash"`
	SecretText string `json:"secretText"`
	CreatedAt string `json:"createdAt"`
	ExpiresAt string `json:"expiresAt"`
	RemainingViews uint64 `json:"remainingViews"`
}
