package bps

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AreaModel struct {
	KodeBps   string `json:"kode_bps"`
	NamaBps   string `json:"nama_bps"`
	KodeDagri string `json:"kode_dagri"`
	NamaDagri string `json:"nama_dagri"`
}

type PostalCodeModel struct {
	KodeBps string `json:"kode_bps"`
	NamaBps string `json:"nama_bps"`
	KodePos string `json:"kode_pos"`
	NamaPos string `json:"nama_pos"`
}

type SearchModel struct {
	Level     string `json:"level"`
	KodeBps   string `json:"kode_bps"`
	NamaBps   string `json:"nama_bps"`
	KodeDagri string `json:"kode_dagri"`
	NamaDagri string `json:"nama_dagri"`
}

const (
	// Province
	bps_province_url            = "https://sig.bps.go.id/rest-drop-down/getwilayah"
	kemendagri_province_url     = "https://sig.bps.go.id/rest-drop-down-dagri/getwilayah"
	bps_bridging_url            = "https://sig.bps.go.id/rest-bridging/getwilayah?level=provinsi&parent=0"
	bps_kemendagri_bridging_url = "https://sig.bps.go.id/rest-bridging-dagri/getwilayah?level=provinsi&parent=0"

	// Regency
	bps_regency_url        = "https://sig.bps.go.id/rest-bridging/getwilayah?level=kabupaten&parent=%s"
	bps_regency_simple_url = "https://sig.bps.go.id/rest-drop-down/getwilayah?level=kabupaten&parent=%s"

	// District
	bps_district_url        = "https://sig.bps.go.id/rest-bridging/getwilayah?level=kecamatan&parent=%s"
	bps_district_simple_url = "https://sig.bps.go.id/rest-drop-down/getwilayah?level=kecamatan&parent=%s"

	// Village
	bps_village_url        = "https://sig.bps.go.id/rest-bridging/getwilayah?level=desa&parent=%s"
	bps_village_simple_url = "https://sig.bps.go.id/rest-drop-down/getwilayah?level=desa&parent=%s"

	// Postal Code
	province_postal_code = "https://sig.bps.go.id/rest-bridging-pos/getwilayah?level=provinsi&parent=0"
	regency_postal_code  = "https://sig.bps.go.id/rest-bridging-pos/getwilayah?level=kabupaten&parent=%s"
	district_postal_code = "https://sig.bps.go.id/rest-bridging-pos/getwilayah?level=kecamatan&parent=%s"
	village_postal_code  = "https://sig.bps.go.id/rest-bridging-pos/getwilayah?level=desa&parent=%s"

	// Historis
	province_historis_url = "https://sig.bps.go.id/rest-historis/getwilayah?level=provinsi&parent=0"
	regency_historis_url  = "https://sig.bps.go.id/rest-historis/getwilayah?level=kabupaten&parent=%s"
	district_historis_url = "https://sig.bps.go.id/rest-historis/getwilayah?level=kecamatan&parent=%s"
	village_historis_url  = "https://sig.bps.go.id/rest-historis/getwilayah?level=desa&parent=%s"

	// Search
	bps_search_area = "https://sig.bps.go.id/rest-cari/getcari?q=%s" // has to be uppercase
)

// SearchArea returns a list of AreaModel that matches the search query
func SearchArea(search string) ([]SearchModel, error) {
	// search should be uppercase
	search = strings.ToUpper(search)

	data, err := getDataFromURL(fmt.Sprintf(bps_search_area, search))
	if err != nil {
		return nil, err
	}

	var areas []SearchModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetProvinces get all provinces
func GetProvinces() ([]AreaModel, error) {
	data, err := getDataFromURL(bps_bridging_url)
	if err != nil {
		return nil, err
	}

	var areas []AreaModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetRegencies get all regencies by province id (kode_bps)
func GetRegencies(provinceId string) ([]AreaModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(bps_regency_url, provinceId))
	if err != nil {
		return nil, err
	}

	var areas []AreaModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetDistricts get all districts by regency id (kode_bps)
func GetDistricts(regencyId string) ([]AreaModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(bps_district_url, regencyId))
	if err != nil {
		return nil, err
	}

	var areas []AreaModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetVillages get all villages by district id (kode_bps)
func GetVillages(districtId string) ([]AreaModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(bps_village_url, districtId))
	if err != nil {
		return nil, err
	}

	var areas []AreaModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetProvincePostalCode get all province postal code
func GetPostalCodeProvinces() ([]PostalCodeModel, error) {
	data, err := getDataFromURL(province_postal_code)
	if err != nil {
		return nil, err
	}

	var areas []PostalCodeModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetRegencyPostalCode get all regency postal code by province id (kode_bps)
func GetPostalCodeRegencies(provinceId string) ([]PostalCodeModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(regency_postal_code, provinceId))
	if err != nil {
		return nil, err
	}

	var areas []PostalCodeModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetDistrictPostalCode get all district postal code by regency id (kode_bps)
func GetPostalCodeDistricts(regencyId string) ([]PostalCodeModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(district_postal_code, regencyId))
	if err != nil {
		return nil, err
	}

	var areas []PostalCodeModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetVillagePostalCode get all village postal code by district id (kode_bps)
func GetPostalCodeVillages(districtId string) ([]PostalCodeModel, error) {
	data, err := getDataFromURL(fmt.Sprintf(village_postal_code, districtId))
	if err != nil {
		return nil, err
	}

	var areas []PostalCodeModel

	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, err
	}
	return areas, nil
}

// GetProvincesHistoris get all provinces historis
func GetHistorisProvinces() ([]map[string]string, error) {
	data, err := getDataFromURL(province_historis_url)
	if err != nil {
		return nil, err
	}

	var historis []map[string]string

	err = json.Unmarshal(data, &historis)
	if err != nil {
		return nil, err
	}
	return historis, nil
}

// GetRegenciesHistoris get all regencies historis by province id (kode_bps)
func GetHistorisRegencies(provinceId string) ([]map[string]string, error) {
	data, err := getDataFromURL(fmt.Sprintf(regency_historis_url, provinceId))
	if err != nil {
		return nil, err
	}

	var historis []map[string]string

	err = json.Unmarshal(data, &historis)
	if err != nil {
		return nil, err
	}
	return historis, nil
}

// GetDistrictsHistoris get all districts historis by regency id (kode_bps)
func GetHistorisDistricts(regencyId string) ([]map[string]string, error) {
	data, err := getDataFromURL(fmt.Sprintf(district_historis_url, regencyId))
	if err != nil {
		return nil, err
	}

	var historis []map[string]string

	err = json.Unmarshal(data, &historis)
	if err != nil {
		return nil, err
	}
	return historis, nil
}

// GetVillagesHistoris get all villages historis by district id (kode_bps)
func GetHistorisVillages(districtId string) ([]map[string]string, error) {
	data, err := getDataFromURL(fmt.Sprintf(village_historis_url, districtId))
	if err != nil {
		return nil, err
	}

	var historis []map[string]string

	err = json.Unmarshal(data, &historis)
	if err != nil {
		return nil, err
	}
	return historis, nil
}
