package main

import "fmt"

func Search(query string) {
	ind := GetIndex(query)
	if len(ind) != 0 {
		for _, rec := range ind {
			fmt.Println("Matching document id:", rec.ID)
			x := GetDoc(rec.ID)
			fmt.Println(x.Content)
		}
	} else {
		fmt.Println("No results")
	}
}
