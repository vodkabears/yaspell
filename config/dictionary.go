package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

// Dictionary implements flag.Value interface and stores regexp patterns of allowed words.
type Dictionary struct {
	File  string
	Words []*regexp.Regexp
}

func (d Dictionary) String() string {
	return d.File
}

// Set reads content of a dictionary and save it into the memory
func (d *Dictionary) Set(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	var re *regexp.Regexp
	var line []byte
	var words []*regexp.Regexp
	b := bufio.NewReader(f)
	for {
		line, _, err = b.ReadLine()
		if err == io.EOF {
			break
		}

		re, err = regexp.Compile(string(line))
		if err != nil {
			return err
		}

		words = append(words, re)
	}

	d.File = file
	d.Words = words

	err = f.Close()
	if err != nil {
		log.Println(err)
	}

	return nil
}
