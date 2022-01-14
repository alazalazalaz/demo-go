package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"strings"
)

type JWSDecodedHeader struct {
	Alg string   `json:"alg"`
	Kid string   `json:"kid"`
	X5c []string `json:"x5c"`
}

type JWSTransactionDecodedPayload struct {
	TransactionId         string `json:"transactionId"`
	OriginalTransactionId string `json:"originalTransactionId"`
	BundleId              string `json:"bundleId"`
	ProductId             string `json:"productId"`
	PurchaseDate          int64  `json:"purchaseDate"`
	OriginalPurchaseDate  int64  `json:"originalPurchaseDate"`
	Quantity              int64  `json:"quantity"`
	Type                  string `json:"type"`
	InAppOwnershipType    string `json:"inAppOwnershipType"`
	SignedDate            int64  `json:"signedDate"`
	RevocationReason      int64  `json:"revocationReason"`
	RevocationDate        int64  `json:"revocationDate"`
}

// APPLE_ROOT_PEM https://www.apple.com/certificateauthority/
const APPLE_ROOT_PEM = `-----BEGIN CERTIFICATE-----
MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwS
QXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcN
MTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBS
b290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49
AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtf
TjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517
IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySr
MA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gA
MGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4
at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM
6BgD56KyKA==
-----END CERTIFICATE-----`

func main() {
	//处理jws串
	jwsString := "eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlFTURDQ0E3YWdBd0lCQWdJUWFQb1BsZHZwU29FSDBsQnJqRFB2OWpBS0JnZ3Foa2pPUFFRREF6QjFNVVF3UWdZRFZRUURERHRCY0hCc1pTQlhiM0pzWkhkcFpHVWdSR1YyWld4dmNHVnlJRkpsYkdGMGFXOXVjeUJEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURUxNQWtHQTFVRUN3d0NSell4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRJeE1EZ3lOVEF5TlRBek5Gb1hEVEl6TURreU5EQXlOVEF6TTFvd2daSXhRREErQmdOVkJBTU1OMUJ5YjJRZ1JVTkRJRTFoWXlCQmNIQWdVM1J2Y21VZ1lXNWtJR2xVZFc1bGN5QlRkRzl5WlNCU1pXTmxhWEIwSUZOcFoyNXBibWN4TERBcUJnTlZCQXNNSTBGd2NHeGxJRmR2Y214a2QybGtaU0JFWlhabGJHOXdaWElnVW1Wc1lYUnBiMjV6TVJNd0VRWURWUVFLREFwQmNIQnNaU0JKYm1NdU1Rc3dDUVlEVlFRR0V3SlZVekJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCT29UY2FQY3BlaXBOTDllUTA2dEN1N3BVY3dkQ1hkTjh2R3FhVWpkNThaOHRMeGlVQzBkQmVBK2V1TVlnZ2gxLzVpQWsrRk14VUZtQTJhMXI0YUNaOFNqZ2dJSU1JSUNCREFNQmdOVkhSTUJBZjhFQWpBQU1COEdBMVVkSXdRWU1CYUFGRDh2bENOUjAxREptaWc5N2JCODVjK2xrR0taTUhBR0NDc0dBUVVGQndFQkJHUXdZakF0QmdnckJnRUZCUWN3QW9ZaGFIUjBjRG92TDJObGNuUnpMbUZ3Y0d4bExtTnZiUzkzZDJSeVp6WXVaR1Z5TURFR0NDc0dBUVVGQnpBQmhpVm9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQXpMWGQzWkhKbk5qQXlNSUlCSGdZRFZSMGdCSUlCRlRDQ0FSRXdnZ0VOQmdvcWhraUc5Mk5rQlFZQk1JSCtNSUhEQmdnckJnRUZCUWNDQWpDQnRneUJzMUpsYkdsaGJtTmxJRzl1SUhSb2FYTWdZMlZ5ZEdsbWFXTmhkR1VnWW5rZ1lXNTVJSEJoY25SNUlHRnpjM1Z0WlhNZ1lXTmpaWEIwWVc1alpTQnZaaUIwYUdVZ2RHaGxiaUJoY0hCc2FXTmhZbXhsSUhOMFlXNWtZWEprSUhSbGNtMXpJR0Z1WkNCamIyNWthWFJwYjI1eklHOW1JSFZ6WlN3Z1kyVnlkR2xtYVdOaGRHVWdjRzlzYVdONUlHRnVaQ0JqWlhKMGFXWnBZMkYwYVc5dUlIQnlZV04wYVdObElITjBZWFJsYldWdWRITXVNRFlHQ0NzR0FRVUZCd0lCRmlwb2RIUndPaTh2ZDNkM0xtRndjR3hsTG1OdmJTOWpaWEowYVdacFkyRjBaV0YxZEdodmNtbDBlUzh3SFFZRFZSME9CQllFRkNPQ21NQnEvLzFMNWltdlZtcVgxb0NZZXFyTU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBUUJnb3Foa2lHOTJOa0Jnc0JCQUlGQURBS0JnZ3Foa2pPUFFRREF3Tm9BREJsQWpFQWw0SkI5R0pIaXhQMm51aWJ5VTFrM3dyaTVwc0dJeFBNRTA1c0ZLcTdoUXV6dmJleUJ1ODJGb3p6eG1ienBvZ29BakJMU0ZsMGRaV0lZbDJlalBWK0RpNWZCbktQdThteW1CUXRvRS9IMmJFUzBxQXM4Yk51ZVUzQ0JqamgxbHduRHNJPSIsIk1JSURGakNDQXB5Z0F3SUJBZ0lVSXNHaFJ3cDBjMm52VTRZU3ljYWZQVGp6Yk5jd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NakV3TXpFM01qQXpOekV3V2hjTk16WXdNekU1TURBd01EQXdXakIxTVVRd1FnWURWUVFERER0QmNIQnNaU0JYYjNKc1pIZHBaR1VnUkdWMlpXeHZjR1Z5SUZKbGJHRjBhVzl1Y3lCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVMTUFrR0ExVUVDd3dDUnpZeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUVic1FLQzk0UHJsV21aWG5YZ3R4emRWSkw4VDBTR1luZ0RSR3BuZ24zTjZQVDhKTUViN0ZEaTRiQm1QaENuWjMvc3E2UEYvY0djS1hXc0w1dk90ZVJoeUo0NXgzQVNQN2NPQithYW85MGZjcHhTdi9FWkZibmlBYk5nWkdoSWhwSW80SDZNSUgzTUJJR0ExVWRFd0VCL3dRSU1BWUJBZjhDQVFBd0h3WURWUjBqQkJnd0ZvQVV1N0Rlb1ZnemlKcWtpcG5ldnIzcnI5ckxKS3N3UmdZSUt3WUJCUVVIQVFFRU9qQTRNRFlHQ0NzR0FRVUZCekFCaGlwb2RIUndPaTh2YjJOemNDNWhjSEJzWlM1amIyMHZiMk56Y0RBekxXRndjR3hsY205dmRHTmhaek13TndZRFZSMGZCREF3TGpBc29DcWdLSVltYUhSMGNEb3ZMMk55YkM1aGNIQnNaUzVqYjIwdllYQndiR1Z5YjI5MFkyRm5NeTVqY213d0hRWURWUjBPQkJZRUZEOHZsQ05SMDFESm1pZzk3YkI4NWMrbGtHS1pNQTRHQTFVZER3RUIvd1FFQXdJQkJqQVFCZ29xaGtpRzkyTmtCZ0lCQkFJRkFEQUtCZ2dxaGtqT1BRUURBd05vQURCbEFqQkFYaFNxNUl5S29nTUNQdHc0OTBCYUI2NzdDYUVHSlh1ZlFCL0VxWkdkNkNTamlDdE9udU1UYlhWWG14eGN4ZmtDTVFEVFNQeGFyWlh2TnJreFUzVGtVTUkzM3l6dkZWVlJUNHd4V0pDOTk0T3NkY1o0K1JHTnNZRHlSNWdtZHIwbkRHZz0iLCJNSUlDUXpDQ0FjbWdBd0lCQWdJSUxjWDhpTkxGUzVVd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NVFF3TkRNd01UZ3hPVEEyV2hjTk16a3dORE13TVRneE9UQTJXakJuTVJzd0dRWURWUVFEREJKQmNIQnNaU0JTYjI5MElFTkJJQzBnUnpNeEpqQWtCZ05WQkFzTUhVRndjR3hsSUVObGNuUnBabWxqWVhScGIyNGdRWFYwYUc5eWFYUjVNUk13RVFZRFZRUUtEQXBCY0hCc1pTQkpibU11TVFzd0NRWURWUVFHRXdKVlV6QjJNQkFHQnlxR1NNNDlBZ0VHQlN1QkJBQWlBMklBQkpqcEx6MUFjcVR0a3lKeWdSTWMzUkNWOGNXalRuSGNGQmJaRHVXbUJTcDNaSHRmVGpqVHV4eEV0WC8xSDdZeVlsM0o2WVJiVHpCUEVWb0EvVmhZREtYMUR5eE5CMGNUZGRxWGw1ZHZNVnp0SzUxN0lEdll1VlRaWHBta09sRUtNYU5DTUVBd0hRWURWUjBPQkJZRUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdEZ1lEVlIwUEFRSC9CQVFEQWdFR01Bb0dDQ3FHU000OUJBTURBMmdBTUdVQ01RQ0Q2Y0hFRmw0YVhUUVkyZTN2OUd3T0FFWkx1Tit5UmhIRkQvM21lb3locG12T3dnUFVuUFdUeG5TNGF0K3FJeFVDTUcxbWloREsxQTNVVDgyTlF6NjBpbU9sTTI3amJkb1h0MlFmeUZNbStZaGlkRGtMRjF2TFVhZ002QmdENTZLeUtBPT0iXX0.eyJ0cmFuc2FjdGlvbklkIjoiMjAwMDA5MDQzMTI3NDUiLCJvcmlnaW5hbFRyYW5zYWN0aW9uSWQiOiIyMDAwMDkwNDMxMjc0NSIsImJ1bmRsZUlkIjoiY29tLnRhcDRmdW4uYnouYXBwc3RvcmUiLCJwcm9kdWN0SWQiOiJiel8wOTk5X2dvbGQiLCJwdXJjaGFzZURhdGUiOjE2MzAwMjQyNzYwMDAsIm9yaWdpbmFsUHVyY2hhc2VEYXRlIjoxNjMwMDI0Mjc2MDAwLCJxdWFudGl0eSI6MSwidHlwZSI6IkNvbnN1bWFibGUiLCJpbkFwcE93bmVyc2hpcFR5cGUiOiJQVVJDSEFTRUQiLCJzaWduZWREYXRlIjoxNjQwOTIyNzAwNjQxLCJyZXZvY2F0aW9uUmVhc29uIjowLCJyZXZvY2F0aW9uRGF0ZSI6MTY0MDkyMTc4ODAwMH0.PLYzUWpjhEihWuMiBgYpU6PoVCTuWtn9KMvNn-nAVSnt2MNwvA7lvisIjkRkRSA0MtOC_WO-pddjeFvp_aPxnA"
	err := handleJWS(jwsString)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("success")
}

func handleJWS(jwsString string) error {
	jwsArray := strings.Split(jwsString, ".")
	if len(jwsArray) != 3 {
		return errors.New("len of jws array != 3")
	}

	//JWS的header包含算法和几个证书
	var jwsHeader JWSDecodedHeader
	jwsHeaderBytes, err := base64.RawStdEncoding.DecodeString(jwsArray[0])
	if err != nil {
		return errors.New(fmt.Sprintf("decode jws header error, err:%v", err))
	}

	if err := json.Unmarshal(jwsHeaderBytes, &jwsHeader); err != nil {
		return errors.New(fmt.Sprintf("unmarshal jws header error, err:%v", err))
	}

	//JWS的payload包含明文
	var jwsPayload JWSTransactionDecodedPayload
	jwsPayloadBytes, err := base64.RawStdEncoding.DecodeString(jwsArray[1])
	if err != nil {
		return errors.New(fmt.Sprintf("decode jws payload error, err:%v", err))
	}

	if err := json.Unmarshal(jwsPayloadBytes, &jwsPayload); err != nil {
		return errors.New(fmt.Sprintf("unmarshal jws payload error, err:%v", err))
	}

	//签名
	//jwsSignBytes, err := base64.RawStdEncoding.DecodeString(jwsArray[2])
	//if err != nil {
	//	return errors.New(fmt.Sprintf("decode jws sign error, err:%v", err))
	//}
	//fmt.Println(string(jwsSignBytes))

	if len(jwsHeader.X5c) <= 0 {
		return errors.New(fmt.Sprintf("len of jwsHeader.X5c <= 0"))
	}

	// 把x5c中的证书字符串转换为标准格式的证书
	pemChainsFormatted := convertPemFormatted(jwsHeader.X5c)

	// 把根证书append到证书链
	pemChainsFormatted = append(pemChainsFormatted, APPLE_ROOT_PEM)

	//验证证书链
	if err := verifyPemChains(pemChainsFormatted); err != nil {
		return err
	}

	//验签数据
	if len(pemChainsFormatted) <= 0 {
		return errors.New(fmt.Sprintf("len of pemChainsFormatted <= 0"))
	}

	// 加载第一个证书
	firstCert, err := loadCertificate(pemChainsFormatted[1])
	if err != nil {
		return errors.New(fmt.Sprintf("loadCertificate failed err:%v", err))
	}

	//验签
	token, err := jwt.Parse(jwsString, func(token *jwt.Token) (interface{}, error) {
		return firstCert.PublicKey, nil
	})

	if err != nil {
		return errors.New(fmt.Sprintf("jwt parse failed, err:%v", err))
	}

	fmt.Println(token)
	//claims, ok := token.Claims.(jwt.MapClaims)
	//if !ok {
	//	return errors.New("token.claim failed")
	//}

	return nil
}

func verifyPemChains(pemChainsFormatted []string) error {
	if len(pemChainsFormatted) <= 1 {
		return errors.New("verifyPemChains failed, len(verifyPemChains)<=1")
	}

	// 循环证书链，证书链中第一个证书是由第二个证书签发的，第二个是由第三个签发的，依次。。。
	// 如果有两个需要验证一次，如果有三个需要验证两次，依次。。。
	for i, _ := range pemChainsFormatted {
		//最后一个证书为跟证书，无需验证
		if i+1 == len(pemChainsFormatted) {
			break
		}

		//加载子证书
		childCert, err := loadCertificate(pemChainsFormatted[i])
		if err != nil {
			return errors.New(fmt.Sprintf("loadCertificate failed num=%d, error:%v", i, err))
		}

		//加载父证书
		parentCert, err := loadCertificate(pemChainsFormatted[i+1])
		if err != nil {
			return errors.New(fmt.Sprintf("loadCertificate failed num=%d, error:%v", i+1, err))
		}

		//验证 子证书 是否由 父证书 签发的
		if err := childCert.CheckSignatureFrom(parentCert); err != nil {
			return errors.New(fmt.Sprintf("check signaturefrom error; err:%v", err))
		}

		fmt.Printf("CheckSignatureFrom success, num=%d \r\n", i)
	}

	return nil
}

func convertPemFormatted(pemChains []string) []string {
	var pemChainsFormatted []string
	for _, pemString := range pemChains {
		var pemRune []rune
		for i, pemSingleRune := range []rune(pemString) {
			pemRune = append(pemRune, pemSingleRune)
			if (i+1)%64 == 0 {
				pemRune = append(pemRune, '\n')
			}
		}
		temp := "-----BEGIN CERTIFICATE-----\n" + string(pemRune) + "\n-----END CERTIFICATE-----"
		fmt.Println(temp)
		pemChainsFormatted = append(pemChainsFormatted, temp)
	}

	return pemChainsFormatted
}

func loadCertificate(pemString string) (*x509.Certificate, error) {
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
