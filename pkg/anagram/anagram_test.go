package anagram

import (
	"reflect"
	"testing"
)

func TestAnagramsAreFound(t *testing.T) {
	input := []string{"abc", "fun", "bac", "fun", "cba", "unf"}

	expected := [][]string{
		{"abc", "bac", "cba"},
		{"fun", "fun", "unf"},
	}

	actual, err := FindAnagrams(input)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v does not equal %v", actual, expected)
	}
}

func TestAnagramsHandlesEmptyInput(t *testing.T) {
	input := []string{}

	expected := [][]string{}

	actual, err := FindAnagrams(input)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v does not equal %v", actual, expected)
	}
}
