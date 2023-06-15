package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(is301Game("G301"))
	fmt.Println(is301Game("g301"))
	fmt.Println(is301Game("g302"))
}

func is301Game(msg string) bool {
	switch msg {
	case "G301":
	case "g301":
		return false
	case "WZ":
	case "wz":
	case "G301_MD":
	case "g301_md":
	case "g301md":
	default:
		return false
	}

	return true
}

func adyenValidateHMAC(hexHmacKey string) bool {
	payload := fmt.Sprintf("%s:%s:%s:%s:%d:%s:%s:%s", "ZWBS24T4FQHG5S82", "", "KingdomGuard", "4ab8fc68470fe50e3f5114e55e056c2f", "785", "HKD", "AUTHORISATION", "true")
	hmacKey, err := hex.DecodeString(hexHmacKey)
	if err != nil {
		fmt.Println(err)
		return false
	}
	mac := hmac.New(sha256.New, hmacKey)
	mac.Write([]byte(payload))
	expectedMAC := mac.Sum(nil)
	messageMAC, err := base64.StdEncoding.DecodeString(`M/Uw+XJd5fCpxpmjLoF35pjExsVKtwcAqr7JtFLdtfA=`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return hmac.Equal(messageMAC, expectedMAC)
}
