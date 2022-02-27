package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/go-scraper/kemdikbud"
)

func main() {
	start := time.Now()

	var provinces []kemdikbud.ProvinceModel
	var districts []kemdikbud.DistrictModel
	var subdistricts []kemdikbud.SubdistrictModel
	var villages []kemdikbud.VillageModel

	var b []byte

	// PROVINCE
	provinces = kemdikbud.GetProvinces()
	log.Printf("################## Province Data: %d ##################", len(provinces))
	b, _ = json.Marshal(provinces)
	log.Println(string(b))

	// DISTRICT
	districts = append(districts, kemdikbud.GetDistricts(provinces[0])...)
	log.Printf("################## District Data: %d ##################", len(districts))
	b, _ = json.Marshal(districts)
	log.Println(string(b))

	// SUBDISTRICT
	subdistricts = append(subdistricts, kemdikbud.GetSubdistricts(districts[0])...)
	log.Printf("################## Subdistrict Data: %d ##################", len(subdistricts))
	b, _ = json.Marshal(subdistricts)
	log.Println(string(b))

	// VILLAGE

	villages = append(villages, kemdikbud.GetVillages(subdistricts[0])...)
	log.Printf("################## Village Data: %d ##################", len(villages))
	b, _ = json.Marshal(villages)
	log.Println(string(b))

	// Elapesd
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Execution", elapsed)
}
