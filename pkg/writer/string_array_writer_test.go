package writer

import (
	"testing"
)

func TestWriterDoesNotError(t *testing.T) {
	input := []string{"abc", "bac", "cba"}
	writer := NewStringSliceFmtWriter(input)

	if err := writer.Write(); err != nil {
		t.Error(err)
	}
}
