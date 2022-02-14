package trigrams

import "testing"

func TestSortingWorks(t *testing.T) {
	var m = map[string]int{
		"a": 9,
		"b": 7,
		"c": 19,
		"d": 11,
		"e": 20,
		"f": 3,
	}

	sorted := Sort(&m)
	if sorted[0].key != "e" || sorted[0].val != 20 {
		t.Fatalf(`SortingIsOk("%v") = '%v'; want '%v'`, sorted, sorted[0], "20 e")
	}
}
