package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func filter(IDs []int, Docs []XMLDocument) (ret []XMLDocument) {
	for _, id := range IDs {
		doc := Docs[id]
		ret = append(ret, doc)
	}
	return
}
func main() {

	if _, err := os.Stat("wiki.gob"); os.IsNotExist(err) {
		fmt.Println("Loading wikipedia...")
		fmt.Println("We only need to do this once...")
		docs := LoadWikipedia()
		file, _ := os.Create("wiki.gob")
		defer file.Close()

		encoder := gob.NewEncoder(file)

		encoder.Encode(docs)
		log.Println("Saved Wikipedia corpus to disk")
	}

	fmt.Println("Loading saved Wikipedia corpus...")
	docs := LoadCorpus()

	if _, err := os.Stat("index.gob"); os.IsNotExist(err) {
		fmt.Println("Creating index...")
		fmt.Println("We only need to do this once...")
		start := time.Now()
		idx := make(Index)
		idx.addToIndex(docs)
		log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
		idx.SaveToDisk()
		log.Println("Saved index to disk")
	}

	idx := LoadIndex()
	fmt.Println("Loaded saved index")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/search", func(c *gin.Context) {
		searchQ := c.Query("q") // shortcut for c.Request.URL.Query().Get("lastname")
		matchedIDs := idx.search(searchQ)
		results := filter(matchedIDs, docs)
		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})
	})
	r.Run()
}

func LoadWikipedia() []XMLDocument {
	start := time.Now()
	docs, err := LoadDocs("enwiki-latest-abstract1.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	return docs
}

func LoadCorpus() []XMLDocument {
	corpus := []XMLDocument{}

	f, err := os.Open("wiki.gob")
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()
	enc := gob.NewDecoder(f)
	if err := enc.Decode(&corpus); err != nil {
		panic("cant decode")
	}
	return corpus
}

func LoadIndex() Index {
	i2 := make(Index)

	f, err := os.Open("index.gob")
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()
	enc := gob.NewDecoder(f)
	if err := enc.Decode(&i2); err != nil {
		panic("cant decode")
	}
	return i2

}
