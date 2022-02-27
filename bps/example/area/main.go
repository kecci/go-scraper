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

	var err error
	var b []byte
	var areaRes []bps.AreaModel
	// var searchRes []bps.SearchModel

	log.Println("################## BPS Province Data: ##################")
	areaRes, err = bps.GetProvinces()
	if err != nil {
		log.Println(err.Error())
	}
	b, _ = json.Marshal(areaRes)
	log.Println(string(b))

	// log.Println("################## BPS Regency Data: ##################")
	// areaRes, err = bps.GetRegencies(areaRes[0].KodeBps)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(areaRes)
	// log.Println(string(b))

	// log.Println("################## BPS District Data: ##################")
	// areaRes, err = bps.GetDistricts(areaRes[0].KodeBps)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(areaRes)
	// log.Println(string(b))

	// log.Println("################## BPS Village Data: ##################")
	// areaRes, err = bps.GetVillages(areaRes[0].KodeBps)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(areaRes)
	// log.Println(string(b))

	// log.Println("################## BPS Search Area Data: ##################")
	// searchRes, err = bps.SearchArea("jakarta")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(searchRes)
	// log.Println(string(b))

	// Elapesd
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Execution", elapsed)
}
