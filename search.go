package main

func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx Index2) search(text string) []int {
	var r []int
	for _, token := range cleanText(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

func concordance(content string) map[string]int {
	// words := strings.Split(content, " ")
	words := filterStems(filterStopWords(filterLowercase(tokenize(content))))
	conc := make(map[string]int, 100)
	for _, word := range words {
		count, exists := conc[word]
		if exists {
			conc[word] = count + 1
		} else {
			conc[word] = 1
		}
	}
	return conc
}
