package main

import (
	"encoding/json"
	"log"

	"github.com/kecci/goscraper/dewanpers"
)

func main() {
	// Max index 1731, page 174
	res, err := dewanpers.GetMedia(1)
	if err != nil {
		log.Println(err.Error())
	}
	b, _ := json.Marshal(res)
	log.Println(string(b))
}
