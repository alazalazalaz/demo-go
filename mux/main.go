package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main(){
	re := mux.NewRouter().StrictSlash(true)
	re.Methods(http.MethodGet).Path("/info").HandlerFunc(InfoGet)
	re.Methods(http.MethodPost).Path("/info").HandlerFunc(InfoPost)

	if err := http.ListenAndServe(":9999", re); err != nil {
		log.Printf("error, err:%v", err)
	}
	log.Println("over")
}

func InfoGet(resp http.ResponseWriter, req *http.Request){
	log.Println("im InfoGet")
}

func InfoPost(resp http.ResponseWriter, req *http.Request){
	log.Println("im POST POST POST")
}