package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type fas struct {
	Word  string
	Count int
}

func Top10(s string) []string {
	r := make([]string, 0, 1)
	if len(s) == 0 {
		return r
	}
	wc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w]++
	}
	f := make([]fas, 0, len(wc))
	for k, v := range wc {
		f = append(f, fas{k, v})
	}
	sort.Slice(f, func(i, j int) bool {
		return f[j].Count < f[i].Count || f[j].Count == f[i].Count && strings.Compare(f[i].Word, f[j].Word) == -1
	})
	for _, w := range f {
		r = append(r, w.Word)
	}
	if len(r) > 10 {
		return r[:10]
	}
	return r
}
