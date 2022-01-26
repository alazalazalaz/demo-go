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
	testA := `-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7aXZrCCnzUQop\nIP2Wv6c7qpvdAt+cKcbZqBVbXVj4tc5Y+CM4Y2y/k+PTNRroySWa7GeRFFEXEo5l\nTWmAvNyojRvb5cqROgClRH7nx3bnnCKu6jLPRsbyQCwrRtF9eWPOWCDwA632pyAQ\nHGDmhglt3bs+bbxoIRfpAQeYW4jpM8Tdd4ZbB9cRih6wLL3pJlJBJQ5eSae7bjSf\n6r12JtpJ6GxUnPLYdf96+NXCGvbT+SL9N1QntoK+eWhMEALzD9nKbzr4DnS7E791\nWTsLXpwP6JJr1PaIJxIw++NK+oAsSrl4k4VMwtuDSuGOnMdwP0M8bbjQcX6ApMpG\nVAjT9lNrAgMBAAECggEAGp3Vd0lXM+Wm4hp7tC+PSmEQ9n2UUR3PI6sZ2rpSNYMU\nmvwlfGUFKobld2H2ar3j57HQL6xJVqkzAgjRhWV37T7v3aDazm9aF3QNQKZ1KQW6\nw0l4YonJyC6nAepykYsY9VL83TmYjmGZTdB0o/4V/r2Gp0zMfPlzkjWsOiOIjp37\nkYuj2h9Fc7L49gLJt1hwJ6ZCjyRR44fWvbppB95pG9fwRiBpGuXS+yCuFXgVPwUL\nzpcbLIo4mPEmJ/LKC7wTPeLVqv9IAtEOsur8gY5btzyhj8bHIBvphwZtfKBEUQbc\n5dihHLHtDcEVh2zxIuFDXAAYxcM7SHKsjx35Yz3JgQKBgQDq4nKqGEtbddD3KG1T\ntQ8hBewm2kh77uVrqt6YM5fUyX3jPoisft06VSjJHD81OtE0CySJTui0ZJ9I8M7Y\nQouuLr0NqrOi18ayWer+0RanSVXHE+i7jWJVXwRI7P0hCyoqCMmcWx6SJ00WXtxD\nUqfYAYmk3MYCmPxldVkzsPuNwQKBgQDMQn8myQkx0xGB3Rv+C0nXX2gKdAmBB4nQ\nQftR0XxkgNrcBNCzI5SXwgcPqPrqy/dwZkpNZph4vC1pZSX+JVLF/+HvwJujTV+o\nSVy0E3nzijqTVWty/+R1ver5QVP02FZCfz/6PDIIJYBzANEfUUwt2guUibIvQm6Z\nHeKfyB6EKwKBgQC9ZmgrvHdXgsND6Xdu3jsa19m64p08QkEA9sGTXVb7IJbAXTZ8\nUbg2R2Eh+gMF+y2il7GYfuaqP9EpUfSY8eAMTmfn01QS5Ye1XUhevV9U+Y+PfgBW\n7AT07i9YVrEuqgAP8RIRYHBgTgydE1TtfIMXbPX+2H9arN4pOdz6D6ZogQKBgGa2\n2bHrTlBbz2X6hmUN0CgZtfvrtgvIeqjbRUIzkdYIZgZV9Yzfy4006NAbLXJ89wTy\nq1KY0PuFxWBGmQBgTGt87WTe+mT4N64slg2H7b9mN2Z3gfzmaUA8LLkC/PEXCOmP\ni+CN/KrUEO/D7WVOQHyeNUyQoaXaRjuBJ8EppQ7ZAoGAGzNigvEgDjgpNz6bbyhk\nmhSxWud1s4fnJ6J2CWfmU/1RBXp96GaeVbqj0Fr5/NFtkZooAOmMvjTVKdns4yHs\nGCouEU3Hv7rde9LSxkb6RiJxxyRTNYlODfk7oZ1W46YtHjvN7+pDYDWbtBXgb2dm\nCbKGDDcWXW/42FGu98HB75k=\n-----END PRIVATE KEY-----\n`
	//testA := `-----BEGIN PRIVATE KEY-----
	//MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQD2EJp0zLzEaC8a
	//VYFwmTPFBtUAZwhr9CptOqHUjEpV+BknkdmMYFEjpmIENEb5SVLqqgWfbMYvpa1U
	//tsbfbxkHNQgggONxNAUIah6jqlVH4LgAvugQEZvKLX2iY4303xhQSJWFA/0KHXsb
	//M3JRxMrM16wFyYsndU4g6jZgbTuUSCm6xUtW79G905PHuH2LaH10dnOovm0jzL5N
	//3sJGNm1jVdbMGTjYjoRF1RlIUMVcvsaKLgRorbxlT2Qbe7WdpiiGZFFaXaz3ivIv
	//M0cuzoiGT9YaAW8rT/zAp7g+PemDlgHrcI0CHhfzWm52vHHBHipiHecT2q+qMsGq
	//JABaQ/ZtAgMBAAECggEAQqkz1uGiYAkjfZxdgT4X3tAfg0jmaKyQk9/W+Kh6+PhK
	//BETIbXWodaJ9Gkt5UM6S9nzSIiMLO9xD7VfI374OxzuQNJIGS+Q+Ws24+DjMerj/
	//0D/nyXajwG17hDM5QXXw9cr8KrluYz+iWregyFMqZ70j3rt/hdSLMnKRAaWGBwDP
	//wdwpJYp52qc6X24HHZgeTQVqILGcQfWPHL4Js6kMRKkQp65tn//gdre52SxEOfVs
	//mjfLESq5nYYPDP+6p+P4t803vG+MXgNG7s6rLeQID/6Mn9SJE6/zBk7POPo0pF69
	//mbcjC96FiuAXudz6q5vyrxGw6IK12IHHfVjbpUlOFwKBgQD93rFyAHvUDxjU7Jb5
	//r1jwMfvWBApkV5xAuUvpZ60wYMU8tcBfeVpouIA6zZXga8LerE5tSb10t25zlRuS
	//UcBOGXw31XPEfKJguKLmtb0vX1PIEHBMgZaAuPoa+dIEeRKAJD0N+KRhXoJYRORg
	//qyzrBHZ6aMvPqeH5yuTcFFsbpwKBgQD4ISUoo33MG8ruK1SFbfRub/P4W7JxGHfK
	//dvZYbiT/0gw3tf+4uoSuTII62v0krxhCc4VwepSuFBlLb2NHaf7VFiRXWjNs9bvO
	//OvRxrlXU1kmn0QOHZ4x5EC41axnkjpPx5IwgFNYnZ0dKdOGytKY288GfuPh6M2g0
	//SwIisS/PywKBgFq7v9HN6HPnHWnkUSVWDURzqN1ZvlkUkuzzDPWTr2tt+rBzCRlW
	//ZwCsiG/70EhQW2p4TON0injvolM/BIasHz4Kj8Ho0SJ1pRdKhjZM2BvZRzlm3qwJ
	//WMAS3JjNlskweHNCAGxA5IdEXvOrU7BVHY21n56qW8WRowlDAb7Cq7lbAoGALTx9
	//tcb4rDl06a7knd7J7UhjxieRhKTdUgAWMVEDVdBJge6gTGFUxyITq/84n3N8jMDa
	//1wLCEGqBogsQonhkiRycS4CaV4cFUjfVUNRjuYXHr9yfNBRgu8GDayG3er+zoSn1
	//kfO6hzyA9sYQT7A3jVUrO59RTZYJrV7vaLu857kCgYAKuM76hLKX1FvtiJlabJFH
	//XHmMEJhr7Z3ZGYEsgI7Qr77PRUBIBgjJhjhBvug/8fDprpWbIi2a1HVFdVKnBZaB
	//e25TbAcep/ocVL0apLlfBSS/d9piJjCAIBrJyuaIfLi5oJhZtp4n/138d4DGDRxj
	//7TMCBtmi2dhraNwZi2jsNg==
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
