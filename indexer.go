package main

import (
	"encoding/gob"
	"os"
)

type Index2 map[string][]int

func (idx Index2) SaveToDisk() {
	file, _ := os.Create("index.gob")
	defer file.Close()

	encoder := gob.NewEncoder(file)

	encoder.Encode(idx)
}

func cleanText(text string) []string {
	return filterStems(filterStopWords(filterLowercase(tokenize(text))))
}

func (idx Index2) addToIndex(docs []XMLDocument) {
	for _, doc := range docs {
		for _, token := range cleanText(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}
