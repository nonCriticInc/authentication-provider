package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/nonCriticInc/authentication-provider/config"
	"github.com/nonCriticInc/authentication-provider/data"
	"gopkg.in/mgo.v2/bson"
)

type CertBuilder interface {
	InitRsaKeyPair() CertBuilder
	OrgId(string) CertBuilder
	Validate() (certBuilder,error)
	Persist() (CertBuilder,error)
	GetCerts() []*data.Cert
}

type certBuilder struct {
	orgId string
	privateKey string
	publicKey string
}


func (c *certBuilder) InitRsaKeyPair() CertBuilder {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	c.privateKey=exportRsaPrivateKeyAsPemStr(privkey)
	c.publicKey, _ =exportRsaPublicKeyAsPemStr(&privkey.PublicKey)
	return c
}
func (c *certBuilder) OrgId(orgId string) CertBuilder{
	c.orgId=orgId
	return c
}


func (c certBuilder) Validate() (certBuilder,error) {
	if(c.orgId==""){
		return c,errors.New("No OrgId Found!")
	}else if(c.publicKey==""){
		return c,errors.New("No Public Key Found!")
	}else if(c.privateKey==""){
		return c,errors.New("No Private Key Found!")
	}
	return c,nil
}

func (c *certBuilder) Persist() (CertBuilder,error) {
	cert:=data.Cert{
		OrgId:        c.orgId,
		PublicKey:    c.publicKey,
		PrivateKey:   c.privateKey,
	}
	err := config.CertManager.Save(&cert)
	if err != nil {
		return c,err
	}
	return c,nil
}


func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func exportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}






func (c certBuilder) GetCerts() []*data.Cert {

	query := bson.M{}
	certList := []*data.Cert{}
	err := config.CertManager.Find(query).Query.All(&certList)
	if err != nil {
		return nil
	}

	return certList
}


func NewCertBuilder() CertBuilder {
	return &certBuilder{}
}
