package main

import (
	"strings"
)

// take a list of docs, clean, save to json, split by spaces, then build a concordance
func index(docs []string) bool {
	for _, doc := range docs {
		id := AddDoc(doc)
		conc := concordance(doc)
		for word, count := range conc {
			ind := GetIndex(word)
			if len(ind) == 0 {
				x := Record{id, count}
				thing := []Record{x}
				AddRecords(word, thing)
			} else {
				x := Record{id, 0}
				ind := []Record{x}
				AddRecords(word, ind)
			}
		}
	}
	return true
}

func concordance(content string) map[string]int {
	words := strings.Split(content, " ")
	conc := make(map[string]int, 100)
	for _, word := range words {
		word = strings.ToLower(word)
		count, exists := conc[word]
		if exists {
			conc[word] = count + 1
		} else {
			conc[word] = 1
		}
	}
	return conc
}

func cleanDoc(doc Document) {

}
