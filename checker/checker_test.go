package checker_test

import (
	"regexp"
	"testing"

	"github.com/VodkaBears/yaspell/checker"
	"github.com/VodkaBears/yaspell/config"
	"github.com/VodkaBears/yaspell/reader"
)

func TestChecherCheck(t *testing.T) {
	cfg := config.NewConfig()
	chunk := &reader.Chunk{"file.txt", "test"}

	if err := checker.Check(chunk, cfg); err != nil {
		t.Error("Expected to successfully check text")
	}
}

func TestChecherCheckSpellError(t *testing.T) {
	cfg := config.NewConfig()
	chunk := &reader.Chunk{"file.txt", "thiswrdisrelybed"}

	if err := checker.Check(chunk, cfg); err == nil {
		t.Error("Expected to get spell errors")
	}
}

func TestChecherCheckJSONParsingError(t *testing.T) {
	cfg := config.NewConfig()
	chunk := &reader.Chunk{"file.txt", "test"}
	cfg.Format = "badformat"

	if err := checker.Check(chunk, cfg); err == nil {
		t.Error("Expected to get JSON parsing error")
	}
}

func TestChecherCheckDictionary(t *testing.T) {
	cfg := config.NewConfig()
	chunk := &reader.Chunk{"file.txt", "thiswrdisrelybed, thiswrdisrelybedtoo"}
	cfg.Dictionary = config.Dictionary{
		File:  "dict.txt",
		Words: []*regexp.Regexp{regexp.MustCompile(`^thiswrdisrelybed$`)},
	}

	if err, ok := checker.Check(chunk, cfg).(checker.Error); ok && len(err.Misspells) > 1 {
		t.Error("Expected to filter misspells with a dictionary")
	}
}
