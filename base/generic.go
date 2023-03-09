package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//关于泛型
	goods := &goodsInfo{
		Name:    "xxx",
		ExtInfo: `{"price":"123", "discount_price":"345"}`,
	}

	re, err := goods.ExtInfo.UnmarshalToGame1()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(re.Price, re.DiscountPrice)
	}
}

type goodsInfo struct {
	Name    string       `json:"name"`
	ExtInfo GoodsExtInfo `json:"ext_info"`
}

type GoodsExtInfo string

type Game1 struct {
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
}

func (g *GoodsExtInfo) UnmarshalToGame1() (*Game1, error) {
	data := &Game1{}
	goodsByte := []byte(*g)
	fmt.Println(*g)
	fmt.Println(goodsByte)
	err := json.Unmarshal(goodsByte, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
