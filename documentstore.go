package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/segmentio/ksuid"
)

type Document struct {
	ID      ksuid.KSUID `json:"ID"`
	Content string      `json:"Content"`
}

// Add a new document to the DocumentStore
func AddDoc(content string) ksuid.KSUID {
	docID := ksuid.New()
	document := Document{docID, content}
	doc, err := json.Marshal(&document)
	if err != nil {
		log.Fatalln(err)
	}
	_ = ioutil.WriteFile(fmt.Sprintf("documents/%v.json", docID), doc, 0644)
	return docID
}

// Get a document by ID
func GetDoc(docID ksuid.KSUID) Document {
	var doc Document
	jsonFile, err := os.Open(fmt.Sprintf("documents/%v.json", docID))
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &doc)
	return doc
}

// Clear json file
func clearDoc() {

}
