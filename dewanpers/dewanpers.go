package dewanpers

import (
	"bytes"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	dewanpers_media_url = "https://datapers.dewanpers.or.id/site/iframe-verified?page=%d"
)

type MediaModel struct {
	Index             string `json:"index"`
	MediaName         string `json:"media_name"`
	MediaType         string `json:"media_type"`
	ResponsiblePerson string `json:"responsible_person"`
	EditorInChief     string `json:"editor_in_chief"`
	Corporation       string `json:"corporation"`
	Province          string `json:"province"`
	Address           string `json:"address"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Website           string `json:"website"`
	Status            string `json:"status"`
	ApprovedDate      string `json:"approved_date"`
}

func GetMedia(page int) ([]MediaModel, error) {
	res, err := getDataFromURL(fmt.Sprintf(dewanpers_media_url, page))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	var media []MediaModel

	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			var headings []string
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			var row []string
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			if len(row) >= 12 && row[1] != "" {
				media = append(media, MediaModel{
					Index:             row[0],
					MediaName:         row[1],
					MediaType:         row[2],
					ResponsiblePerson: row[3],
					EditorInChief:     row[4],
					Corporation:       row[5],
					Province:          row[6],
					Address:           row[7],
					Phone:             row[8],
					Email:             row[9],
					Website:           row[10],
					Status:            row[11],
					ApprovedDate:      row[12],
				})
			}
		})
	})

	return media, nil
}
