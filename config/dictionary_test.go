package config_test

import (
	"testing"

	"github.com/yaspell/config"
)

func TestDictionaryString(t *testing.T) {
	file := "file.txt"
	d := config.Dictionary{File: file}

	if d.String() != file {
		t.Errorf("Expected %s, got %s", file, d)
	}
}

func TestDictionarySet(t *testing.T) {
	var d config.Dictionary
	file := "testdata/dict.txt"

	err := d.Set(file)
	if err != nil {
		t.Log("Unexpected error:", err)
	}

	if d.File != file {
		t.Errorf("Expected to set %s as a filename, had %s", file, d.File)
	}

	secondRe := d.Words[1].String()
	expectedRe := "^(G|g)oroutines$"
	if secondRe != expectedRe {
		t.Errorf("Expected to have %s as a second regexp, had %s", secondRe, expectedRe)
	}
}

func TestDictionarySetOpenError(t *testing.T) {
	var d config.Dictionary
	err := d.Set("blabla")

	if err == nil {
		t.Error("Expected to get an error from os.Open")
	}
}

func TestDictionarySetRegexpError(t *testing.T) {
	var d config.Dictionary
	err := d.Set("testdata/incorrect-dict.txt")

	if err == nil {
		t.Error("Expected to get regexp compile error")
	}
}
