package main

import (
	"encoding/xml"
	"os"

	"github.com/cheggaaa/pb/v3"
)

type XMLDocument struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocs(path string) ([]XMLDocument, error) {
	var limit int64 = 1024 * 1024 * 950
	bar := pb.Full.Start64(limit)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	barReader := bar.NewProxyReader(f)
	decoder := xml.NewDecoder(barReader)
	dump := struct {
		Documents []XMLDocument `xml:"doc"`
	}{}
	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}

	bar.Finish()
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
