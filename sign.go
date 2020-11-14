package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func Sha1WithRsa(data, privateKey []byte) (sign string, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("privateKey error")
		return
	}

	var private *rsa.PrivateKey
	if private, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}

	t := crypto.SHA1.New()
	if _, err = t.Write(data); err != nil {
		return
	}

	digest := t.Sum(nil)

	var signedData []byte
	if signedData, err = rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, digest); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signedData)
	return
}
