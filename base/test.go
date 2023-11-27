package main

import (
	"log"
)

type ContentState map[string]interface{}

func main() {
	t := "abcdefghijklmnopqrstuvwxyz1"
	prefix := t[len(t)-16:]
	log.Printf(prefix)
}

func getNeed() []string {
	return nil
}
