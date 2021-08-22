package writer

import (
	"fmt"
	"strings"
)

// StringSliceFmtWriter writes a string slice with fmt
type StringSliceFmtWriter struct {
	Strings []string
}

// NewStringSliceFmtWriter writes a string slice with fmt
func NewStringSliceFmtWriter(strings []string) Writer {
	return &StringSliceFmtWriter{
		Strings: strings,
	}
}

// Write a string slice with fmt, separated by commas
func (writer *StringSliceFmtWriter) Write() error {
	fmt.Println(strings.Join(writer.Strings, ","))
	return nil
}
