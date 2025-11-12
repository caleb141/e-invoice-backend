package utility

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
)

type CryptoKeys struct {
	PublicKey   *rsa.PublicKey
	Certificate string
}

func LoadCryptoKeys(filename string) (*CryptoKeys, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", filename, err)
	}

	type keyData struct {
		PublicKey   string `json:"public_key"`
		Certificate string `json:"certificate"`
	}

	var kd keyData
	if err := json.Unmarshal(content, &kd); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	pemData, err := base64.StdEncoding.DecodeString(kd.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 public_key: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block for public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return &CryptoKeys{
		PublicKey:   rsaPub,
		Certificate: kd.Certificate,
	}, nil
}
