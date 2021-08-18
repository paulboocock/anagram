package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:      "Anagram",
		Usage:     "Finds anagrams from a word list",
		Version:   "0.1.0",
		Copyright: "(c) 2021 Paul Boocock",
		Compiled:  time.Now(),
		Authors: []cli.Author{
			{
				Name: "Paul Boocock",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
