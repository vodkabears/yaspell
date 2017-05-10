package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yaspell/checker"
	"github.com/yaspell/config"
	"github.com/yaspell/reader"
)

func main() {
	cfg := config.NewConfig()
	flag.StringVar(&cfg.Lang, "lang", cfg.Lang, "Language to check")
	flag.StringVar(&cfg.Format, "format", cfg.Format, "Text format")
	flag.Var(&cfg.Dictionary, "dict", "Dictionary with regexp patterns")
	flag.Var(&cfg.Options, "opts", "Yaspeller options")
	flag.Var(&cfg.Version, "v", "Prints current version")
	flag.Usage = func() {
		fmt.Printf("Usage of %s [FLAGS...] [FILES...]:\n", os.Args[0])
		flag.PrintDefaults()
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
