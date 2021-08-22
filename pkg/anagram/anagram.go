package anagram

import (
	"sort"

	"github.com/paulboocock/anagram/pkg/sorting"
)

// FindAnagrams when given a slice of strings, this will return a slice where
// each element contains a slice with words which are anagrams
func FindAnagrams(values []string) ([][]string, error) {
	anagrams := map[string][]string{}

	for _, value := range values {
		valueAsRunes := []rune(value)
		sort.Sort(sorting.SortAlphabetically(valueAsRunes))
		sortedValue := string(valueAsRunes)

		anagrams[sortedValue] = append(anagrams[sortedValue], value)
	}

	result := [][]string{}
	for _, value := range anagrams {
		if len(value) > 1 {
			result = append(result, value)
		}
	}

	return result, nil
}
