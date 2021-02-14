# Tinydex

Mini full-text search engine written in Go

Currently using a small portion of the [Wikipedia English abstract corpus](https://dumps.wikimedia.org/enwiki/latest/)

### Features

- [x] Search results appear in microseconds
- [x] Tokenizer
- [ ] Boolean queries (AND, OR, NOT)
- [x] Save index to disk
- [ ] Simple result ranking with concordance
- [ ] Multithreaded indexing
- [x] API

### How to use

1. Download the latest XML dump of the Wikipedia English abstract corpus (for example: `enwiki-latest-abstract1.xml.gz`)
2. Unzip
3. `go run main.go cleanup.go indexer.go loadDocs.go search.go`
4. Wait a little while for your computer to load the entire Wikipedia corpus (this only happens once)
5. Search with `localhost:8080/search?q=spinach`, where q = your search query
6. Wikipedia search results appear in microseconds

You can also curl it if you really want to.
```
$ curl "http://localhost:8080/search?q=beans"
```

### How it currently works:

1. Download 600,000+ documents from the Wikipedia abstract corpus (~950 MB)
2. Tokenize + stem + clean up text
3. Build inverted index, save it to disk in binary format
4. Standard search stuff happens here
