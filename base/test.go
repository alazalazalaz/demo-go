package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	log.Println(IsVailedIp("101.24.90.164"))
}

type exchange struct {
	Rates map[string]float64 `json:"rates"`
}

func IsVailedIp(ip string) bool {

	data := `{
    "disclaimer": "Usage subject to terms: https://openexchangerates.org/terms",
    "license": "https://openexchangerates.org/license",
    "timestamp": 1662001200,
    "base": "USD",
    "rates": {
      "AED": 3.67299,
      "AFN": 87.797968,
      "ALL": 116.837856,
      "AMD": 401.91949,
      "ANG": 1.793968,
      "AOA": 429.42085,
      "ARS": 138.7215,
      "AUD": 1.46755,
      "AWG": 1.805,
      "AZN": 1.7,
      "BAM": 1.951176,
      "BBD": 2,
      "BDT": 94.589829,
      "BGN": 1.950047,
      "BHD": 0.377022,
      "BIF": 2053.430871,
      "BMD": 1,
      "BND": 1.391526,
      "BOB": 6.878295,
      "BRL": 5.1835,
      "BSD": 1,
      "BTC": 0.000049757873,
      "BTN": 79.13004,
      "BWP": 12.794102,
      "BYN": 2.512433,
      "BZD": 2.006375,
      "CAD": 1.316205,
      "CDF": 2043.518542,
      "CHF": 0.980221,
      "CLF": 0.032508,
      "CLP": 896.879022,
      "CNH": 6.913265,
      "CNY": 6.9034,
      "COP": 4402.846156,
      "CRC": 637.07819,
      "CUC": 1,
      "CUP": 25.75,
      "CVE": 110.1,
      "CZK": 24.4597,
      "DJF": 177.203796,
      "DKK": 7.4205,
      "DOP": 52.925093,
      "DZD": 140.434559,
      "EGP": 19.2413,
      "ERN": 15,
      "ETB": 52.534456,
      "EUR": 0.997656,
      "FJD": 2.2382,
      "FKP": 0.862969,
      "GBP": 0.862969,
      "GEL": 2.89,
      "GGP": 0.862969,
      "GHS": 10.003342,
      "GIP": 0.862969,
      "GMD": 55.15,
      "GNF": 8593.704096,
      "GTQ": 7.70391,
      "GYD": 208.242911,
      "HKD": 7.848215,
      "HNL": 24.497083,
      "HRK": 7.4973,
      "HTG": 115.464142,
      "HUF": 399.921834,
      "IDR": 14877.377407,
      "ILS": 3.327175,
      "IMP": 0.862969,
      "INR": 79.554944,
      "IQD": 1452.763054,
      "IRR": 42350,
      "ISK": 141.38,
      "JEP": 0.862969,
      "JMD": 149.928913,
      "JOD": 0.709,
      "JPY": 139.48916667,
      "KES": 119.548882,
      "KGS": 80.699291,
      "KHR": 4086.906232,
      "KMF": 490.224587,
      "KPW": 900,
      "KRW": 1352.62519,
      "KWD": 0.30825,
      "KYD": 0.829497,
      "KZT": 470.862467,
      "LAK": 15352.209429,
      "LBP": 1505.065807,
      "LKR": 353.347833,
      "LRD": 153.800043,
      "LSL": 16.991167,
      "LYD": 4.913184,
      "MAD": 10.539862,
      "MDL": 19.210905,
      "MGA": 4186.881226,
      "MKD": 61.468195,
      "MMK": 2090.295942,
      "MNT": 3206.445782,
      "MOP": 8.046368,
      "MRU": 37.61,
      "MUR": 44.54594,
      "MVR": 15.36,
      "MWK": 1021.664489,
      "MXN": 20.213202,
      "MYR": 4.484,
      "MZN": 63.899991,
      "NAD": 16.96,
      "NGN": 418.065024,
      "NIO": 35.773252,
      "NOK": 9.992,
      "NPR": 126.61549,
      "NZD": 1.640003,
      "OMR": 0.385011,
      "PAB": 1,
      "PEN": 3.798978,
      "PGK": 3.50733,
      "PHP": 56.387497,
      "PKR": 218.191464,
      "PLN": 4.716783,
      "PYG": 6857.356053,
      "QAR": 3.641,
      "RON": 4.8315,
      "RSD": 117.035991,
      "RUB": 60.125004,
      "RWF": 1039.178451,
      "SAR": 3.759171,
      "SBD": 8.244163,
      "SCR": 12.871598,
      "SDG": 579,
      "SEK": 10.695095,
      "SGD": 1.400222,
      "SHP": 0.862969,
      "SLL": 13748.9,
      "SOS": 565.834028,
      "SRD": 25.126,
      "SSP": 130.26,
      "STD": 22392.090504,
      "STN": 24.45,
      "SVC": 8.709965,
      "SYP": 2512.53,
      "SZL": 16.99366,
      "THB": 36.6775,
      "TJS": 10.137512,
      "TMT": 3.5,
      "TND": 3.195,
      "TOP": 2.357826,
      "TRY": 18.2036,
      "TTD": 6.747804,
      "TWD": 30.48,
      "TZS": 2321.231316,
      "UAH": 36.744563,
      "UGX": 3796.378676,
      "USD": 1,
      "UYU": 40.78914,
      "UZS": 10879.681804,
      "VES": 7.84425,
      "VND": 23462.5,
      "VUV": 116.90004,
      "WST": 2.689325,
      "XAF": 654.41913,
      "XAG": 0.05643883,
      "XAU": 0.00058663,
      "XCD": 2.70255,
      "XDR": 0.738968,
      "XOF": 654.41913,
      "XPD": 0.00048044,
      "XPF": 119.051973,
      "XPT": 0.00118485,
      "YER": 250.249998,
      "ZAR": 17.175951,
      "ZMW": 15.752017,
      "ZWL": 322
    }
  }`

	exchangeData := exchange{}
	err := json.Unmarshal([]byte(data), &exchangeData)
	if err != nil {
		log.Printf("err:%v", err)
		return false
	}

	//for k, _ := range exchangeData.Rates {
	//fmt.Println(k)
	//}

	// adyenData
	adyenData := `AFA
ALL
ARP
ATS
AUD
BBD
BEF
BGL
BHD
BIF
BOP
BRC
BSD
BUK
CAD
CHF
CLP
CNY
COP
CRC
CSK
CUP
CYP
DEM
DJF
DKK
DOP
DZD
ECS
EEK
EGP
ESP
EUR
FIM
FJD
FRF
GBP
GHC
GMD
GNS
GQE
GRD
GTQ
GWP
GYD
HKD
HNL
HTG
HUF
IDR
IEP
INR/Rs/Re
IQD
IRR
ISK
ITL
JMD
JOD
JPY
KES
KHR
KMF
KPW/KRW
KWD
LAK
LBP
LKR
LRD
LTL
LUF
LVL
LYD
MAD
MCF
MOP
MRO
MTP
MUR
MVR
MXP
MYR
NGN
NIC
NLG
NOK
NPR
NZD
OMR
PAB
PES
PHP
PLZ
PRK
PT
PYG
QAR
ROL
RU
RWF
SAR
SBD
SCR
SDP
SEK
SGD
SLL
SOS
SRG
SVC
SYP
THP
TND
TRL
TTD
TZS
UAH
UGS
USD
UYP
VEB
VND
XAF
XAF
XAF
XAF
XAF
XOF
XOF
XOF
XOF
XOF
XOF
YDD
YER
YUD
ZAR
ZMK
ZRZ
ZWD`

	//存在于adyen不存在与web site的货币
	noneCurrencyMap := make([]string, 0)
	adyenMap := strings.Split(adyenData, "\n")
	for _, v := range adyenMap {
		_, exist := exchangeData.Rates[v]
		if !exist {
			noneCurrencyMap = append(noneCurrencyMap, v)
		}
	}

	fmt.Printf("len:%v \r\n", len(noneCurrencyMap))
	fmt.Println(noneCurrencyMap)

	return true
}
