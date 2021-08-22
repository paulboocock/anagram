package reader

import (
	"bufio"
	"os"
)

type StringBatchByLengthReaderConfig struct {
	Filepath string
}

type StringBatchByLengthReader struct {
	CurrentLength int

	Config  *StringBatchByLengthReaderConfig
	File    *os.File
	Scanner *bufio.Scanner
}

func NewStringBatchByLengthReader(filepath string) Reader {
	return &StringBatchByLengthReader{
		Config: &StringBatchByLengthReaderConfig{
			Filepath: filepath,
		},
	}
}

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

func (reader *StringBatchByLengthReader) Close() error {
	return reader.File.Close()
}
