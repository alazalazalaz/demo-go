package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
)

func main() {
	testA := `-----BEGIN PRIVATE KEY-----\nxxx\n-----END PRIVATE KEY-----\n`
	//testA := `-----BEGIN PRIVATE KEY-----
	//xxx
	//-----END PRIVATE KEY-----`

	cer, err := ParseKey([]byte(testA))
	if err != nil {
		log.Printf("error:%v", err)
	}

	log.Println(cer)
}

func loadCertificateFromPemString(pemString string) (*x509.Certificate, error) {
	blocker, _ := pem.Decode([]byte(pemString))
	if blocker == nil {
		return nil, errors.New("first pem.Decode error, blocker is nil")
	}

	cert, err := x509.ParseCertificate(blocker.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("first ParseCertificate error,err:%v", err))
	}

	return cert, nil
}

// ParseKey 将pem格式的私钥，转换为*rsa.PrivateKey，当然也可以转为其他算法的key，比如ecdsa,dsa等。
func ParseKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return parsed, nil
}
