# go-scraper

Go library to scrap or get dataset.

## Example
Installation:
```sh
go get github.com/kecci/go-scraper/bps
```

Implementation:
```go
package main

import (
	"encoding/json"
	"log"

	"github.com/kecci/go-scraper/bps"
)

func main() {
	areaRes, err := bps.GetProvinces()
	if err != nil {
		log.Println(err.Error())
	}
	b, _ := json.Marshal(areaRes)
	log.Println(string(b))
}
```

Result:
```json
[{"kode_bps":"11","nama_bps":"ACEH","kode_dagri":"11","nama_dagri":"ACEH"},{"kode_bps":"12","nama_bps":"SUMATERA UTARA","kode_dagri":"12","nama_dagri":"SUMATERA UTARA"},{"kode_bps":"13","nama_bps":"SUMATERA BARAT","kode_dagri":"13","nama_dagri":"SUMATERA BARAT"},{"kode_bps":"14","nama_bps":"RIAU","kode_dagri":"14","nama_dagri":"RIAU"},{"kode_bps":"15","nama_bps":"JAMBI","kode_dagri":"15","nama_dagri":"JAMBI"},{"kode_bps":"16","nama_bps":"SUMATERA SELATAN","kode_dagri":"16","nama_dagri":"SUMATERA SELATAN"},{"kode_bps":"17","nama_bps":"BENGKULU","kode_dagri":"17","nama_dagri":"BENGKULU"},{"kode_bps":"18","nama_bps":"LAMPUNG","kode_dagri":"18","nama_dagri":"LAMPUNG"},{"kode_bps":"19","nama_bps":"KEP. BANGKA BELITUNG","kode_dagri":"19","nama_dagri":"KEP. BANGKA BELITUNG"},{"kode_bps":"21","nama_bps":"KEP. RIAU","kode_dagri":"21","nama_dagri":"KEP. RIAU"},{"kode_bps":"31","nama_bps":"DKI JAKARTA","kode_dagri":"31","nama_dagri":"DKI JAKARTA"},{"kode_bps":"32","nama_bps":"JAWA BARAT","kode_dagri":"32","nama_dagri":"JAWA BARAT"},{"kode_bps":"33","nama_bps":"JAWA TENGAH","kode_dagri":"33","nama_dagri":"JAWA TENGAH"},{"kode_bps":"34","nama_bps":"DI YOGYAKARTA","kode_dagri":"34","nama_dagri":"DI YOGYAKARTA"},{"kode_bps":"35","nama_bps":"JAWA TIMUR","kode_dagri":"35","nama_dagri":"JAWA TIMUR"},{"kode_bps":"36","nama_bps":"BANTEN","kode_dagri":"36","nama_dagri":"BANTEN"},{"kode_bps":"51","nama_bps":"BALI","kode_dagri":"51","nama_dagri":"BALI"},{"kode_bps":"52","nama_bps":"NUSA TENGGARA BARAT","kode_dagri":"52","nama_dagri":"NUSA TENGGARA BARAT"},{"kode_bps":"53","nama_bps":"NUSA TENGGARA TIMUR","kode_dagri":"53","nama_dagri":"NUSA TENGGARA TIMUR"},{"kode_bps":"61","nama_bps":"KALIMANTAN BARAT","kode_dagri":"61","nama_dagri":"KALIMANTAN BARAT"},{"kode_bps":"62","nama_bps":"KALIMANTAN TENGAH","kode_dagri":"62","nama_dagri":"KALIMANTAN TENGAH"},{"kode_bps":"63","nama_bps":"KALIMANTAN SELATAN","kode_dagri":"63","nama_dagri":"KALIMANTAN SELATAN"},{"kode_bps":"64","nama_bps":"KALIMANTAN TIMUR","kode_dagri":"64","nama_dagri":"KALIMANTAN TIMUR"},{"kode_bps":"65","nama_bps":"KALIMANTAN UTARA","kode_dagri":"65","nama_dagri":"KALIMANTAN UTARA"},{"kode_bps":"71","nama_bps":"SULAWESI UTARA","kode_dagri":"71","nama_dagri":"SULAWESI UTARA"},{"kode_bps":"72","nama_bps":"SULAWESI TENGAH","kode_dagri":"72","nama_dagri":"SULAWESI TENGAH"},{"kode_bps":"73","nama_bps":"SULAWESI SELATAN","kode_dagri":"73","nama_dagri":"SULAWESI SELATAN"},{"kode_bps":"74","nama_bps":"SULAWESI TENGGARA","kode_dagri":"74","nama_dagri":"SULAWESI TENGGARA"},{"kode_bps":"75","nama_bps":"GORONTALO","kode_dagri":"75","nama_dagri":"GORONTALO"},{"kode_bps":"76","nama_bps":"SULAWESI BARAT","kode_dagri":"76","nama_dagri":"SULAWESI BARAT"},{"kode_bps":"81","nama_bps":"MALUKU","kode_dagri":"81","nama_dagri":"MALUKU"},{"kode_bps":"82","nama_bps":"MALUKU UTARA","kode_dagri":"82","nama_dagri":"MALUKU UTARA"},{"kode_bps":"91","nama_bps":"PAPUA BARAT","kode_dagri":"92","nama_dagri":"PAPUA BARAT"},{"kode_bps":"94","nama_bps":"PAPUA","kode_dagri":"91","nama_dagri":"PAPUA"}]
```


## Scraper List

| No | Media | Type | dataset | source | website |
|---|---|---|---|---|---|
| 1 | bps | government | province, regency, district, village, postal code | REST API | https://sig.bps.go.id/ |
| 2 | kemdikbud | government | province, regency, district, village, postal code | Web Scraping | https://www.kemdikbud.go.id/ |
| 3 | kemdagri | government | province, regency, district, village | CSV | https://www.kemendagri.go.id/ |
| 4 | dewanpers | government | news portal, media, redaction, radio | Web Scraping | https://dewanpers.or.id/ |
| 5 | gadm | organization | geojson | JSON | https://gadm.org/ |
| 6 | instagram | social  media | profile, post | REST API | https://www.instagram.com/ |