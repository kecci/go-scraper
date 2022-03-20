package youtube

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/kecci/goscraper"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	// Youtube search url
	search_url = "https://www.youtube.com/results?search_query=%s"
)

func Search(search string) (interface{}, error) {
	res, err := goscraper.GetDataFromURL(fmt.Sprintf(search_url, search))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(doc.Html())

	return nil, nil
}

// GetVideoInfo to scrapping video info from youtube.com
// example url: https://www.youtube.com/watch?v=c5vpmV6AuRI
func GetVideoInfo(url string) (interface{}, error) {
	res, err := goscraper.GetDataFromURL(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body").Each(func(i int, item *goquery.Selection) {
		attr, exists := item.Attr("dir")
		if exists && attr == "ltr" {
			log.Println(item.Text()[:1000])
		}
	})

	doc.Find("meta").Each(func(index int, item *goquery.Selection) {
		attr, exists := item.Attr("property")
		if exists && attr == "og:url" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:title" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:description" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:image" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:video" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:video:url" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:video:type" {
			log.Println(item.Attr("content"))
		}

		if exists && attr == "og:video:tag" {
			log.Println(item.Attr("content"))
		}
	})

	return nil, nil
}

func GetYoutubeSerice(apiKey string) *youtube.Service {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	return youtubeService
}

func ChannelsListByUsername(service *youtube.Service, part []string, forUsername string) {
	call := service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

func ChannelsListByChannelId(service *youtube.Service, part []string, channelId string) {
	call := service.Channels.List(part)
	call = call.Id(channelId)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
