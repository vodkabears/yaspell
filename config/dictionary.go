package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

// Dictionary stores regexp patterns of allowed words
type Dictionary struct {
	File  string
	Words []*regexp.Regexp
}

func (d Dictionary) String() string {
	return d.File
}

// Set implements interface of flag.Value (https://golang.org/pkg/flag/#Value)
func (d *Dictionary) Set(file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	b := bufio.NewReader(f)
	for {
		line, _, err := b.ReadLine()
		if err == io.EOF {
			break
		}

		r, err := regexp.Compile(string(line))
		if err != nil {
			log.Fatal(err)
		}

		d.Words = append(d.Words, r)
	}

	d.File = file

	return nil
}
