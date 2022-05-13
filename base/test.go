package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func main() {
	log.Println(IsVailedIp("101.24.90.164"))
}

func IsVailedIp(ip string) bool {
	b64SignedString := "rhburtTcYWoZIvda72kzxH6fzqWapiFjFhkz1IRKRIMC7fwE7ck3w2Z1pxX20gikLU4KrDorVTRAPTIWYx30+iaGgNyRJ7tAPVOpVPB+d/HObRkpfz0EyvqdvGG9D7s2yc5U2mosGdVLjA8Ylgd0/X1VYkPHGDz14sJwp9hRRGA="
	fmt.Println(b64SignedString)
	b64SignedString = strings.Replace(b64SignedString, " ", "", -1)
	b64SignedString = strings.Replace(b64SignedString, "/", "", -1)
	b64SignedString = strings.Replace(b64SignedString, " ", "", -1)

	fmt.Println(b64SignedString)
	_, err := base64.StdEncoding.DecodeString(b64SignedString)
	_, err2 := base64.URLEncoding.DecodeString(b64SignedString)
	_, err3 := base64.RawStdEncoding.DecodeString(b64SignedString)
	_, err4 := base64.RawURLEncoding.DecodeString(b64SignedString)

	fmt.Println(err)
	fmt.Println(err2)
	fmt.Println(err3)
	fmt.Println(err4)

	return true
}
