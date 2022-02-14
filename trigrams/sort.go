package trigrams

import (
	"sort"
)

// In order to sort a map[string]int, the best approach is to create a struct of key and value properties
// and implement the Golang Sort interface
type sequence struct {
	val int
	key string
}

type sequences []sequence

// Implementation of the Golang Sort Interface.
// It provides primitives for sorting slices and user-defined collections. See https://pkg.go.dev/sort.
func (s sequences) Len() int           { return len(s) }
func (s sequences) Less(i, j int) bool { return s[i].val < s[j].val }
func (s sequences) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Sort sorts in memory a map of strings keys to ints values. It returns the sorted word sequences from a
// specific text with their occurrences.
func Sort(m *map[string]int) sequences {

	var s sequences
	for k, v := range *m {
		s = append(s, sequence{val: v, key: k})
	}

	sort.Sort(sort.Reverse(s))

	return s
}
