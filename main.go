package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/VodkaBears/yaspell/checker"
	"github.com/VodkaBears/yaspell/config"
	"github.com/VodkaBears/yaspell/reader"
)

const helpText = `Usage: yaspell [flags] [files ...]

Yaspell checks spelling of large texts with Yandex.Speller API.

Flags:

-opts
	Yaspeller options.

	Example: -opts=IGNORE_UPPERCASE,IGNORE_DIGITS

	IGNORE_UPPERCASE ignores uppercased words
	IGNORE_DIGITS ignores words with digits
	IGNORE_URLS ignores urls, emails, filenames
	FIND_REPEAT_WORDS highlights repetitions of words, consecutive
	IGNORE_LATIN disables suggestions for incorrect words
	FLAG_LATIN marks latin words as incorrect
	BY_WORDS ignores dictionary context
	IGNORE_CAPITALIZATION ignores the incorrect use of UPPERCASE/lowercase letters
	IGNORE_ROMAN_NUMERALS ignores roman numerals

-dict
	Dictionary file with regexp patterns.

	Example: -dict=dict.txt
	Regexp syntax: https://golang.org/pkg/regexp/syntax/#hdr-Syntax

	dict.txt content:
	^nananana$
	^(?i)gogogogo$

-lang
	Language to check.

	Values: en, ru, uk
	Default: ru,en
	Example: -lang=en,ru,uk

-format
	Text format.

	Values: html, plain
	Default: plain
	Example: -format=html

-version
	Prints current version.
`

func main() {
	cfg := config.NewConfig()
	flag.StringVar(&cfg.Lang, "lang", cfg.Lang, "Language to check")
	flag.StringVar(&cfg.Format, "format", cfg.Format, "Text format")
	flag.Var(&cfg.Dictionary, "dict", "Dictionary file with regexp patterns")
	flag.Var(&cfg.Options, "opts", "Yaspeller options")
	flag.Var(&cfg.Version, "version", "Prints current version")
	flag.Usage = func() {
		fmt.Println(helpText)
	}

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var isError bool
	ch := make(chan reader.Chunk)
	go reader.Read(ch, files)
	for chunk := range ch {
		if err := checker.Check(&chunk, cfg); err != nil {
			isError = true
			fmt.Println(err)
		}
	}

	if isError {
		os.Exit(1)
	}
}
