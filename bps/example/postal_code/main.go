package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/go-scraper/bps"
)

func main() {
	// start time for elapsed
	start := time.Now()

	var postalRes []bps.PostalCodeModel
	var b []byte
	var err error

	log.Println("################## BPS Postal Code Province Data: ##################")
	postalRes, err = bps.GetPostalCodeProvinces()
	if err != nil {
		log.Println(err.Error())
	}
	b, _ = json.Marshal(postalRes)
	log.Println(string(b))

	log.Println("################## BPS Postal Code Regency Data: ##################")
	postalRes, err = bps.GetPostalCodeRegenciesByProvinceId(postalRes[0].KodeBps)
	if err != nil {
		log.Println(err.Error())
	}
	b, _ = json.Marshal(postalRes)
	log.Println(string(b))

	log.Println("################## BPS Postal Code District Data: ##################")
	postalRes, err = bps.GetPostalCodeDistrictsByRegencyId(postalRes[0].KodeBps)
	if err != nil {
		log.Println(err.Error())
	}
	b, _ = json.Marshal(postalRes)
	log.Println(string(b))

	log.Println("################## BPS Postal Code Village Data: ##################")
	postalRes, err = bps.GetPostalCodeVillagesByDistrictId(postalRes[0].KodeBps)
	if err != nil {
		log.Println(err.Error())
	}
	b, _ = json.Marshal(postalRes)
	log.Println(string(b))

	// Elapesd
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Execution", elapsed)
}
