package gadm

import (
	"encoding/json"

	"github.com/kecci/goscraper"
)

const (
	geojson_regency_url = "https://raw.githubusercontent.com/rifani/geojson-political-indonesia/master/IDN_adm_2_kabkota.json"
)

type GeojsonModel struct {
	Bbox     []float64 `json:"bbox"`
	Type     string    `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			ID0       int         `json:"ID_0"`
			Iso       string      `json:"ISO"`
			Name0     string      `json:"NAME_0"`
			ID1       int         `json:"ID_1"`
			Name1     string      `json:"NAME_1"`
			ID2       int         `json:"ID_2"`
			Name2     string      `json:"NAME_2"`
			Varname2  interface{} `json:"VARNAME_2"`
			NlName2   interface{} `json:"NL_NAME_2"`
			Hasc2     string      `json:"HASC_2"`
			Cc2       interface{} `json:"CC_2"`
			Type2     string      `json:"TYPE_2"`
			Engtype2  string      `json:"ENGTYPE_2"`
			Validfr2  string      `json:"VALIDFR_2"`
			Validto2  string      `json:"VALIDTO_2"`
			Remarks2  interface{} `json:"REMARKS_2"`
			ShapeLeng float64     `json:"Shape_Leng"`
			ShapeArea float64     `json:"Shape_Area"`
		} `json:"properties"`
		Geometry struct {
			Type        string            `json:"type"`
			Coordinates [][][]interface{} `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func GetGeojsonRegencies() (*GeojsonModel, error) {
	geojson := GeojsonModel{}
	res, err := goscraper.GetDataFromURL(geojson_regency_url)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &geojson)
	return &geojson, nil
}
