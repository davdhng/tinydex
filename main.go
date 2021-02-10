package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type IndexCollection []struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Timegate string `json:"timegate"`
	CdxAPI   string `json:"cdx-api"`
}

func main() {
	resp, err := http.Get("https://index.commoncrawl.org/collinfo.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var cdxIndexes IndexCollection

	if err := json.NewDecoder(resp.Body).Decode(&cdxIndexes); err != nil {
		log.Fatalln(err)
	}

	latestIndex := cdxIndexes[0].CdxAPI
	fmt.Println(latestIndex)
}
