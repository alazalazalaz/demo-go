package main

import "fmt"

type Car struct {
	Type     string
	Name     string
	WheelNum int
	Light    bool
}

type Bike struct {
	Type string
	Name string
}

func main() {
	var traffic interface{}
	traffic = getMyTraffic(4)
	switch traffic.(type) {
	case *Car:
		fmt.Println("拿到车了")
		break
	case *Bike:
		fmt.Println("拿到自行车了")
		break
	}
	fmt.Println(traffic)
}

func getMyTraffic(wheelNum int) interface{} {
	myCar := &Car{
		Type:     "car",
		Name:     "my car",
		WheelNum: 4,
	}

	myBike := &Bike{
		Type: "bike",
		Name: "my bike",
	}
	if wheelNum > 2 {
		return myCar
	} else {
		return myBike
	}
}
