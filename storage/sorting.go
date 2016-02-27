package storage

import "sort"

// words sorter
type SortWordsByCount func(p1, p2 *Word) bool

// words sorting function
func (by SortWordsByCount) Sort(words []*Word) {
	ws := &wordsSorter{
		words: words,
		by:    by,
	}
	sort.Sort(ws)
}

type wordsSorter struct {
	words []*Word
	by    func(p1, p2 *Word) bool
}

func (s *wordsSorter) Len() int {
	return len(s.words)
}

func (s *wordsSorter) Swap(i, j int) {
	s.words[i], s.words[j] = s.words[j], s.words[i]
}

func (s *wordsSorter) Less(i, j int) bool {
	return s.by(s.words[i], s.words[j])
}

func bycount(p1, p2 *Word) bool {
	return p1.Count > p2.Count
}
