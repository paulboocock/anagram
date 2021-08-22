package sorting

import (
	"sort"
	"testing"
)

func TestSortAlphabetically(t *testing.T) {
	initial := "paul"
	expected := "alpu"

	actualRunes := []rune(initial)
	sort.Sort(SortAlphabetically(actualRunes))
	actual := string(actualRunes)

	if actual != expected {
		t.Errorf("%v does not equal %s", actual, "alpu")
	}
}
