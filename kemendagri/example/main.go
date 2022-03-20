package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/goscraper/kemendagri"
)

func main() {
	// start time for elapsed
	start := time.Now()

	// prepare byte
	var b []byte

	// // Get all provinces / provinsi
	// log.Println("################## Province Data: ##################")
	// provinces := kemendagri.GetProvinces()
	// b, _ = json.Marshal(provinces)
	// log.Println(string(b))

	// // Get all regencies / kabupaten
	// log.Println("################## Regency Data: ##################")
	// regencies := kemendagri.GetRegencies()
	// b, _ = json.Marshal(regencies)
	// log.Println(string(b))

	// Get all districts / kecamatan
	// log.Println("################## District Data: ##################")
	// districts := kemendagri.GetDistricts()
	// b, _ = json.Marshal(districts)
	// log.Println(string(b))

	// Get all villages / kelurahan
	log.Println("################## Village Data: ##################")
	villages := kemendagri.GetVillages()
	b, _ = json.Marshal(villages)
	log.Println(string(b))

	// Elapesd
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Execution", elapsed)
}
