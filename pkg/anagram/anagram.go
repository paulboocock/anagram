package anagram

import (
	"sort"

	"github.com/paulboocock/anagram/pkg/sorting"
)

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
		result = append(result, value)
	}

	return result, nil
}
