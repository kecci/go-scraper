package kemdikbud

import (
	"bytes"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kecci/goscraper"
)

const (
	kemdikbud_host         = "https://referensi.data.kemdikbud.go.id"
	kemdikbud_province_url = "https://referensi.data.kemdikbud.go.id/index11.php"
)

type ProvinceModel struct {
	ProvinceName string `json:"province_name"`
	ProvinceURL  string `json:"province_url"`
}

type DistrictModel struct {
	Province     ProvinceModel `json:"province"`
	DistrictName string        `json:"district_name"`
	DistrictURL  string        `json:"district_url"`
}

type SubdistrictModel struct {
	Province        ProvinceModel `json:"province"`
	District        DistrictModel `json:"district"`
	SubdistrictName string        `json:"subdistrict_name"`
	SubdistrictURL  string        `json:"subdistrict_url"`
}

type VillageModel struct {
	Province    ProvinceModel    `json:"province"`
	District    DistrictModel    `json:"district"`
	Subdistrict SubdistrictModel `json:"subdistrict"`
	VillageName string           `json:"village_name"`
	VillageURL  string           `json:"village_url"`
	PostalCode  string           `json:"postal_code"`
}

// GetProvinces to scrapping Province data from kemebdikbud.go.id
func GetProvinces() (provinces []ProvinceModel) {
	// PROVINCE
	res, err := goscraper.GetDataFromURL(kemdikbud_province_url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("tbody").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(ProvinceModel)
		row.ProvinceName = strings.Title(sel.Find("a").Text())
		row.ProvinceURL, _ = sel.Find("a").Attr("href")
		row.ProvinceURL = kemdikbud_host + "/" + row.ProvinceURL

		if row.ProvinceName != "" && strings.Contains(row.ProvinceURL, "index11.php?") {
			provinces = append(provinces, *row)
		}
	})
	return provinces
}

// GetDistricts to scrapping District data from kemebdikbud.go.id
func GetDistricts(province ProvinceModel) (districts []DistrictModel) {
	// DISTRICT
	res, err := goscraper.GetDataFromURL(province.ProvinceURL)
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody").Children().Each(func(j int, sel *goquery.Selection) {
		row := new(DistrictModel)
		row.Province = province
		row.DistrictName = strings.Title(sel.Find("a").Text())
		row.DistrictURL, _ = sel.Find("a").Attr("href")
		row.DistrictURL = kemdikbud_host + "/" + row.DistrictURL

		if row.DistrictName != "" && strings.Contains(row.DistrictURL, "index11.php?") {
			districts = append(districts, *row)
		}
	})

	return districts
}

// GetSubdistricts to scrapping Subdistrict data from kemebdikbud.go.id
func GetSubdistricts(district DistrictModel) (subdistricts []SubdistrictModel) {
	res, err := goscraper.GetDataFromURL(district.DistrictURL)
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tbody").Children().Each(func(j int, sel *goquery.Selection) {
		row := new(SubdistrictModel)
		row.Province = district.Province
		row.District = district
		row.SubdistrictName = strings.Title(sel.Find("a").Text())
		row.SubdistrictURL, _ = sel.Find("a").Attr("href")
		row.SubdistrictURL = kemdikbud_host + "/" + row.SubdistrictURL

		if row.SubdistrictName != "" && strings.Contains(row.SubdistrictURL, "index11.php?") {
			subdistricts = append(subdistricts, *row)
		}
	})

	return subdistricts
}

// GetVillages to scrapping Village data from kemebdikbud.go.id
func GetVillages(subdistrict SubdistrictModel) (villages []VillageModel) {
	res, err := goscraper.GetDataFromURL(subdistrict.SubdistrictURL)
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".display").Children().Each(func(j int, sel *goquery.Selection) {
		village := new(VillageModel)
		village.Province = subdistrict.Province
		village.District = subdistrict.District
		village.Subdistrict = subdistrict
		village.VillageURL, _ = sel.Find("a").Attr("href")
		village.VillageURL = kemdikbud_host + "/" + village.VillageURL

		if strings.Contains(village.VillageURL, "tabs.php?npsn=") {
			// PostalCode
			res, err = goscraper.GetDataFromURL(village.VillageURL)
			if err != nil {
				log.Fatal(err)
			}

			docPostalCode, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
			if err != nil {
				log.Fatal(err)
			}
			docPostalCode.Find("#tabs-1").Find("table").First().Find("tbody").First().Find("table").First().Find("tbody").First().Children().Each(func(k int, selTbody *goquery.Selection) {
				if k == 3 {
					selTbody.Find("td").First().Remove()
					selTbody.Find("td").First().Remove()
					selTbody.Find("td").First().Remove()
					village.PostalCode = selTbody.Find("td").First().Text()
				}
				if k == 5 {
					selTbody.Find("td").First().Remove()
					selTbody.Find("td").First().Remove()
					selTbody.Find("td").First().Remove()
					village.VillageName = selTbody.Find("td").First().Text()
				}
			})

			villages = append(villages, *village)
		}
	})

	return villages
}
