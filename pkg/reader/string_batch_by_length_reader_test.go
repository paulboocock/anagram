package reader

import (
	"reflect"
	"testing"
)

func TestReaderFirstSetFromDataFile(t *testing.T) {
	expected := []string{"abc", "fun", "bac", "fun", "cba", "unf"}
	fileReader := NewStringBatchByLengthReader("../../data/example1.txt")

	defer fileReader.Close()

	if err := fileReader.Open(); err != nil {
		t.Error(err)
	}

	actual, err := fileReader.Read()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v does not equal %v", actual, expected)
	}
}

func TestReaderBothSetsFromDataFile(t *testing.T) {
	expected1 := []string{"abc", "fun", "bac", "fun", "cba", "unf"}
	expected2 := []string{"hello"}
	fileReader := NewStringBatchByLengthReader("../../data/example1.txt")

	defer fileReader.Close()

	if err := fileReader.Open(); err != nil {
		t.Error(err)
	}

	actual1, err := fileReader.Read()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual1, expected1) {
		t.Errorf("%v does not equal %v", actual1, expected1)
	}

	actual2, err := fileReader.Read()

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual2, expected2) {
		t.Errorf("%v does not equal %v", actual2, expected2)
	}
}

func TestReaderWhenReadingBeyondEOF(t *testing.T) {
	fileReader := NewStringBatchByLengthReader("../../data/example1.txt")

	defer fileReader.Close()

	if err := fileReader.Open(); err != nil {
		t.Error(err)
	}

	fileReader.Read()
	fileReader.Read()
	actual, err := fileReader.Read()

	if err != nil {
		t.Error(err)
	}

	if len(actual) != 0 {
		t.Errorf("Expected Read to produce empty array but it did not: %v", actual)
	}
}

func TestLoadMissingDataFile(t *testing.T) {
	fileReader := NewStringBatchByLengthReader("../../data/missing.txt")

	defer fileReader.Close()

	if err := fileReader.Open(); err == nil {
		t.Errorf("Expected file should be missing")
	}
}
