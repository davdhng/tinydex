package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/segmentio/ksuid"
)

type Record struct {
	ID    ksuid.KSUID `json:"ID"`
	Count int         `json:"Count"`
}

type Index struct {
	Name    string   `json:"Name"`
	Records []Record `json:"Records"`
}

// A document is a JSON object which contains zero or more fields, or key-value pairs. In many cases, a field in a document can be the content of a webpage.
// This is a record-level index. Each word/key is mapped to a list of document references
// word_1  = [document_references]
// ....
// word_n = [document_references]

// Each row in the index has a document id (20 byte ksuid)
// Size on disk to store a million documents
// ((20 x 1,000,000 x 500) / 1024) / 1024 = 9536 Megabytes

// Add records to an index
func AddRecords(name string, records []Record) bool {
	var idx Index
	for _, rec := range records {
		idx.Records = append(idx.Records, rec)
	}
	idx.Name = name
	file, _ := json.Marshal(&idx)

	fName := fmt.Sprintf("idx/%v.json", name)
	err := ioutil.WriteFile(fName, file, 0644)
	if err != nil {
		return false
	}
	return true
}

// Given the name of an index, retrieve all the documents
func GetIndex(name string) []Record {
	var idx Index
	jsonFile, err := os.Open(fmt.Sprintf("idx/%v.json", name))
	if err != nil {
		// fmt.Println(err)
		return idx.Records
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &idx)
	return idx.Records
}

func ClearIndex() {

}
