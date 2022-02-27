package kemendagri

import (
	"encoding/csv"
	"log"
	"os"
)

type Province struct {
	Id   string `csv:"id" json:"id"`
	Name string `csv:"name" json:"name"`
}

type Regency struct {
	Id         string `csv:"id" json:"id"`
	ProvinceId string `csv:"province_id" json:"province_id"`
	Name       string `csv:"name" json:"name"`
}

type District struct {
	Id        string `csv:"id" json:"id"`
	RegencyId string `csv:"regency_id" json:"regency_id"`
	Name      string `csv:"name" json:"name"`
}

type Village struct {
	Id         string `csv:"id" json:"id"`
	DistrictId string `csv:"district_id" json:"district_id"`
	Name       string `csv:"name" json:"name"`
}

// GetProvinces returns all provinces as a slice of Province structs
func GetProvinces() []*Province {
	provinceFile, err := os.OpenFile("kemendagri/csv/provinces.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer provinceFile.Close()

	csvReader := csv.NewReader(provinceFile)
	csvReader.Comma = ';'

	// Read first line to skip header
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	provinces := []*Province{}
	for i := range data {
		provinces = append(provinces, &Province{
			Id:   data[i][0],
			Name: data[i][1],
		})
	}

	return provinces
}

// GetRegencies returns all regencies as a slice of Regency structs
func GetRegencies() []*Regency {
	regenciesFile, err := os.OpenFile("kemendagri/csv/regencies.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer regenciesFile.Close()

	csvReader := csv.NewReader(regenciesFile)
	csvReader.Comma = ';'

	// Read first line to skip header
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	regencies := []*Regency{}
	for i := range data {
		regencies = append(regencies, &Regency{
			Id:         data[i][0],
			ProvinceId: data[i][1],
			Name:       data[i][2],
		})
	}

	return regencies
}

// GetDistricts returns all districts as a slice of District structs
func GetDistricts() []*District {
	districtsFile, err := os.OpenFile("kemendagri/csv/districts.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer districtsFile.Close()

	csvReader := csv.NewReader(districtsFile)
	csvReader.Comma = ';'

	// Read first line to skip header
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	districts := []*District{}
	for i := range data {
		districts = append(districts, &District{
			Id:        data[i][0],
			RegencyId: data[i][1],
			Name:      data[i][2],
		})
	}

	return districts
}

// GetVillages returns all villages as a slice of Village structs
func GetVillages() []*Village {
	villagesFile, err := os.OpenFile("kemendagri/csv/villages.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer villagesFile.Close()

	csvReader := csv.NewReader(villagesFile)
	csvReader.Comma = ';'

	// Read first line to skip header
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	villages := []*Village{}
	for i := range data {
		villages = append(villages, &Village{
			Id:         data[i][0],
			DistrictId: data[i][1],
			Name:       data[i][2],
		})
	}

	return villages
}
