package main

import "github.com/kecci/go-scraper/youtube"

func main() {
	service := youtube.GetYoutubeSerice("AIzaSyAxRRlCIJvGPDZjyWAuoxzLKQsv5B3o-Ck")
	youtube.ChannelsListByUsername(service, []string{"snippet", "contentDetails", "statistics"}, "GoogleDevelopers")
	youtube.ChannelsListByChannelId(service, []string{"snippet", "contentDetails", "statistics"}, "UCvwXwDAxhDqbrfnVQakHpHw")
}
