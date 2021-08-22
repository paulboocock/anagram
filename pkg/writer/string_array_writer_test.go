package writer

import (
	"testing"
)

func TestWriterDoesNotError(t *testing.T) {
	input := []string{"abc", "bac", "cba"}
	writer := NewStringArrayFmtWriter(input)

	if err := writer.Write(); err != nil {
		t.Error(err)
	}
}
