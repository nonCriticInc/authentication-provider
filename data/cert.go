package data

import (
	"github.com/go-bongo/bongo"
)
// Scheme of saving Certificates
type Cert struct {
	bongo.DocumentBase `bson:",inline"`
	OrgId              string    `bson:"orgId"`
	PublicKey              string    `bson:"publicKey"`
	PrivateKey           string    `bson:"privateKey"`
}

