package main

import (
	"fmt"
	"math/rand"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type IndexCollection []struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Timegate string `json:"timegate"`
	CdxAPI   string `json:"cdx-api"`
}

type IndexResult struct {
	URLKey       string `json:"urlkey"`
	Timestamp    string `json:"timestamp"`
	URL          string `json:"url"`
	Mime         string `json:"mime"`
	MimeDetected string `json:"mime-detected"`
	Status       string `json:"status"`
	Digest       string `json:"digest"`
	Length       string `json:"length"`
	Offset       string `json:"offset"`
	Filename     string `json:"filename"`
	CharSet      string `json:"charset"`
	Languages    string `json:"languages"`
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13",
	"Mozilla/5.0 (Windows; U; Windows NT 6.0; de) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13",
	"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; pt-PT; rv:1.9.2.7) Gecko/20100713 Firefox/3.6.7 (.NET CLR 3.5.30729)",
	"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.6) Gecko/20100628 Ubuntu/10.04 (lucid) Firefox/3.6.6 GTB7.1",
	"Mozilla/5.0 (X11; Mageia; Linux x86_64; rv:10.0.9) Gecko/20100101 Firefox/10.0.9",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:9.0a2) Gecko/20111101 Firefox/9.0a2",
	"Mozilla/5.0 (Windows NT 6.2; rv:9.0.1) Gecko/20100101 Firefox/9.0.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
}

func randomOption(options []string) string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(options)
	return options[randNum]
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {

	poop := "Poopy di scoop Scoop diddy whoop Whoop di scoop di poop Poop di scoopty Scoopty whoop Whoopity scoop whoop poop Poop diddy whoop scoop Poop, poop Scoop diddy whoop Whoop diddy scoop Whoop diddy scoop, poop"
	a := "Of course the actual implementation can get pretty hairy."
	b := "The reason for this is that a search is actually more complex then it first appears."
	c := "Take for example the following search term cool stuff. When someone searches for cool stuff they are expecting that a list of documents containing the words cool and stuff will appear in a list."
	poo := []string{poop, a, b, c}
	// concordance(poop)
	index(poo)

	fmt.Println("Done indexing...")
	Search("search")

	// resp, err := http.Get("https://index.commoncrawl.org/collinfo.json")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	// var cdxIndexes IndexCollection

	// if err := json.NewDecoder(resp.Body).Decode(&cdxIndexes); err != nil {
	// 	log.Fatalln(err)
	// }

	// latestIndex := cdxIndexes[0].CdxAPI

	// req, err := http.NewRequest("GET", latestIndex, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// q := url.Values{}
	// q.Add("url", "reddit.com")
	// q.Add("limit", "10")
	// q.Add("output", "json")

	// req.URL.RawQuery = q.Encode()

	// resp2, err := http.Get(req.URL.String())
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp2.Body.Close()
	// x := []IndexResult{}

	// decoder := json.NewDecoder(resp2.Body)
	// bar := pb.StartNew(10)

	// for decoder.More() {
	// 	var result IndexResult
	// 	if err := decoder.Decode(&result); err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	x = append(x, result)
	// 	bar.Increment()
	// }

	// record := x[0]
	// fmt.Println(record)
	// prefixURL := "https://commoncrawl.s3.amazonaws.com/"
	// dataURL := prefixURL + record.Filename
	// startByte, err := strconv.Atoi(record.Offset)
	// recordLength, err := strconv.Atoi(record.Length)
	// endByte := startByte + recordLength + 1

	// req2, err := http.NewRequest("GET", dataURL, nil)
	// req2.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", startByte, endByte))
	// req.Header.Set("User-Agent", randomOption(userAgents))

	// resp3, err := http.DefaultClient.Do(req2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r, err := gzip.NewReader(resp3.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// io.Copy(os.Stdout, r)
	// r.Close()
}
