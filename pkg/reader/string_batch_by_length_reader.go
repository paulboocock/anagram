package reader

import (
	"bufio"
	"os"
)

// StringBatchByLengthReaderConfig object for creating a StringBatchByLengthReader
type StringBatchByLengthReaderConfig struct {
	Filepath string
}

// StringBatchByLengthReader reads strings from a file of sorted strings until the string length changes
type StringBatchByLengthReader struct {
	CurrentLength int

	Config  *StringBatchByLengthReaderConfig
	File    *os.File
	Scanner *bufio.Scanner
}

// NewStringBatchByLengthReader creates a new StringBatchByLengthReader
// Reads strings from a file of sorted strings until the string length changes
func NewStringBatchByLengthReader(filepath string) Reader {
	return &StringBatchByLengthReader{
		Config: &StringBatchByLengthReaderConfig{
			Filepath: filepath,
		},
	}
}

// Open the file for reading
func (reader *StringBatchByLengthReader) Open() error {
	file, err := os.Open(reader.Config.Filepath)
	if err != nil {
		return err
	}
	reader.File = file
	reader.Scanner = bufio.NewScanner(file)

	// Perform initial Scan(), ready for first Read() call
	reader.Scanner.Scan()

	if err := reader.Scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Read the strings from the file until the length of the strings change
func (reader *StringBatchByLengthReader) Read() ([]string, error) {
	var result []string

	// Initial line should be available as Load() calls Scan()
	if reader.Scanner.Text() != "" {
		reader.CurrentLength = len(reader.Scanner.Text())
		result = append(result, reader.Scanner.Text())
	}

	for reader.Scanner.Scan() {
		if len(reader.Scanner.Text()) != reader.CurrentLength {
			break
		}
		result = append(result, reader.Scanner.Text())
	}

	if err := reader.Scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// Close the open file
func (reader *StringBatchByLengthReader) Close() error {
	return reader.File.Close()
}
