# Tinydex

Mini search engine written in Go

Currently using a small portion of the [Wikipedia English abstract corpus](https://dumps.wikimedia.org/enwiki/latest/)

### Features

- [X] Search results appear in microseconds
- [X] Tokenizer
- [ ] Boolean queries (AND, OR, NOT)
- [X] Save index to disk
- [ ] Simple result ranking with concordance
- [ ] Multithreaded indexing
- [ ] API???

### How it currently works:

1. Download 600,000+ documents from the Wikipedia abstract corpus (~950 MB)
2. Tokenize + stem + clean up text
3. Build inverted index, save it to disk in binary format 
4. Standard search stuff happens here 
