package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/paulboocock/anagram/pkg/anagram"
	"github.com/paulboocock/anagram/pkg/reader"
	"github.com/paulboocock/anagram/pkg/writer"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:      "Anagram",
		Usage:     "Finds anagrams from a word list - ./anagram /path/to/file.txt",
		Version:   "0.1.0",
		Copyright: "(c) 2021 Paul Boocock",
		Compiled:  time.Now(),
		Authors: []cli.Author{
			{
				Name: "Paul Boocock",
			},
		},
		Action: func(c *cli.Context) error {
			filepath := c.Args().Get(0)

			if filepath == "" {
				return fmt.Errorf("missing file path argument, application must be run as 'anagram /path/to/file.txt'")
			}

			fmt.Printf("Running with file: %q\n", filepath)

			fileReader := reader.NewStringBatchByLengthReader(filepath)
			defer fileReader.Close()
			if err := fileReader.Open(); err != nil {
				return err
			}

			for currentValues, err := fileReader.Read(); len(currentValues) != 0; currentValues, err = fileReader.Read() {
				if err != nil {
					return err
				}

				anagrams, err := anagram.FindAnagrams(currentValues)

				if err != nil {
					return err
				}

				if len(anagrams) > 0 {
					fmt.Println("----------------------------")
					fmt.Printf("Anagrams with %d characters:\n", len(anagrams[0][0]))
					for _, anagram := range anagrams {
						writer.NewStringSliceFmtWriter(anagram).Write()
					}
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
