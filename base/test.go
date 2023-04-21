package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
)

func main() {
	log.Println(uuid.New().String())
	log.Println(uuid.New().String())
	log.Println(uuid.New().String())
	log.Println(uuid.New().String())
	log.Println(uuid.New().String())
	log.Println(uuid.New().String())
	id := uuid.New().String()
	spanId := id[24:]
	log.Println(id)
	log.Println(spanId)

	log.Println(adyenValidateHMAC("F199834FA0537AC7782B862A32232354CA1BC646F1085C8C654744F702CF3526"))
	//pointtest()

	fmt.Println(strings.SplitN("aa:bb:cc", ":", 2)[0])
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
