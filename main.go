package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Loading wikipedia...")
	docs := LoadWikipedia()
	if _, err := os.Stat("index.gob"); os.IsNotExist(err) {
		start := time.Now()
		idx := make(Index2)
		idx.addToIndex(docs)
		log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
		idx.SaveToDisk()
		log.Println("Saved index to disk")
	}
	idx := LoadIndex()
	fmt.Println("Loaded saved index")
	start := time.Now()
	matchedIDs := idx.search("noodles")
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

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

func LoadIndex() Index2 {
	i2 := make(Index2)

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
