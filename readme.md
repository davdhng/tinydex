# Tinydex

Mini full-text search engine written in Go

Currently using a small portion of the [Wikipedia English abstract corpus](https://dumps.wikimedia.org/enwiki/latest/)

### Features

- [x] Search results appear in <20 ms
- [x] Tokenizer
- [x] Save index to disk
- [x] API
- [ ] Typo tolerance / Fuzzy search 
- [ ] Boolean queries (AND, OR, NOT)
- [ ] Simple result ranking with concordance
- [ ] Multithreaded indexing
- [ ] Faceted search
- [ ] Docker support

### How to use

1. Download & unzip the latest XML dump of the Wikipedia English abstract corpus (for example: `enwiki-latest-abstract1.xml.gz`)
2. `go run main.go cleanup.go indexer.go loadDocs.go search.go`
3. Wait a little while for your computer to load the entire Wikipedia corpus (this only happens once)
4. Search with `localhost:8080/search?q=spinach`, where q = your search query
5. Wikipedia search results appear in milliseconds

You can also curl it if you really want to.

```
$ curl "http://localhost:8080/search?q=beans"
```

### How it currently works:

1. Download 600,000+ documents from the Wikipedia abstract corpus (~950 MB)
2. Tokenize + stem + clean up text
3. Build inverted index, save it to disk in binary format
4. Standard search stuff happens here

### Why you might use it

- Get bulk data on a topic in Wikipedia, like "spacecraft"
- Build a fast beautiful frontend for Wiki search
- Do some kind of data analysis idk
