package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kecci/goscraper/instagram"
)

func main() {
	// start time for elapsed
	start := time.Now()

	username := "keccikun"
	limit := uint16(10)

	// Get account info
	account, err := instagram.GetAccountByUsername(username)
	if err != nil {
		log.Println("instagram.GetAccountByUsername", err.Error())
	}
	if account.ID != "" {
		b, _ := json.Marshal(account)
		log.Println(string(b))
	}

	// Get media info
	// media, err := instagram.GetMediaByCode("code")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(media)
	// log.Println(string(b))

	// Get media by url
	// media, err := instagram.GetMediaByURL("https://instagram.com/p/code")
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(media)
	// log.Println(string(b))

	// Get slice of account media
	media, err := instagram.GetAccountMedia(username, limit)
	if err != nil {
		log.Println("instagram.GetAccountMedia", err.Error())
	}
	if len(media) > 0 {
		b, _ := json.Marshal(media)
		log.Println(string(b))
	}

	// or slice of all account media
	media, err = instagram.GetAllAccountMedia(username)
	if err != nil {
		log.Println("instagram.GetAllAccountMedia", err.Error())
	}
	if len(media) > 0 {
		b, _ := json.Marshal(media)
		log.Println(string(b))
	}

	// Get slice of location last media
	// media, err = instagram.GetLocationMedia("location_id", limit)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(media)
	// log.Println(string(b))

	// Get array[9] of location top media
	// media, err := instagram.GetLocationTopMedia("location_id")

	// Get location info
	// location, err := instagram.GetLocationByID("location_id")

	// Get slice of tag last media
	// media, err = instagram.GetTagMedia("tag", limit)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// b, _ = json.Marshal(media)
	// log.Println(string(b))

	// Get array[9] of tag top media
	mediaTop, err := instagram.GetLocationTopMedia("tag")
	if err != nil {
		log.Println("instagram.GetLocationTopMedia", err.Error())
	}
	if len(mediaTop) > 0 {
		b, _ := json.Marshal(mediaTop)
		log.Println(string(b))
	}

	// Search for users (return slice of Accounts)
	users, err := instagram.SearchForUsers(username)
	if err != nil {
		log.Println("instagram.SearchForUsers", err.Error())
	}
	if len(users) > 0 {
		b, _ := json.Marshal(users)
		log.Println(string(b))
	}

	// Elapesd
	elapsed := time.Since(start)
	log.Printf("%s took %s", "Execution", elapsed)
}
