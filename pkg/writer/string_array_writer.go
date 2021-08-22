package writer

import (
	"fmt"
	"strings"
)

type StringArrayFmtWriter struct {
	Strings []string
}

func NewStringArrayFmtWriter(strings []string) Writer {
	return &StringArrayFmtWriter{
		Strings: strings,
	}
}

func (writer *StringArrayFmtWriter) Write() error {
	fmt.Println(strings.Join(writer.Strings, ","))
	return nil
}
