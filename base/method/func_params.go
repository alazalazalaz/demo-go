



package main 

import "fmt"

func main(){

	host := "127.0.0.1"

	connect(host)
	connect2(host, "root", true)

	// fmt.Printf("xx\n")
}


// func connect(host string, user string, pw){
func connect(host string, other ...string){
	user, pw := "", ""
	if len(other) == 1 {
		user = other[0]
	}else if len(other) == 2 {
		user = other[0]
		pw = other[1]
	}

	fmt.Printf("%s, %s, %s\n", host, user, pw)
}

func connect2(host string, other ...interface{}){
	user := other[0]
	pw := other[1]
	fmt.Printf("%s, %s, %v\n", host, user, pw)
}

