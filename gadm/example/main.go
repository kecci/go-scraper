package main

import (
	"encoding/json"
	"log"

	"github.com/kecci/goscraper/gadm"
)

func main() {
	regencies, err := gadm.GetGeojsonRegencies()
	if err != nil {
		log.Println(err.Error())
	}
	b, _ := json.Marshal(regencies.Features[4])
	log.Println(string(b))
}
