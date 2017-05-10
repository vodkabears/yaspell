package config

import (
	"fmt"
	"strings"
)

const (
	// IgnoreUppercase ignores "uppercased" words
	IgnoreUppercase = 1
	// IgnoreDigits ignores words with digits
	IgnoreDigits = 2 << 0
	// IgnoreUrls ignores urls, emails, filenames
	IgnoreUrls = 2 << 1
	// FindRepeatWords highlights repetitions of words, consecutive
	FindRepeatWords = 2 << 2
	// IgnoreLatin ignores latin words
	IgnoreLatin = 2 << 3
	// NoSuggest disables suggestions for incorrect words
	NoSuggest = 2 << 4
	// FlagLatin marks latin words as incorrect
	FlagLatin = 2 << 5
	// ByWords ignores a dictionary context
	ByWords = 2 << 6
	// IgnoreCapitalization ignores the incorrect use of UPPERCASE/lowercase letters
	IgnoreCapitalization = 2 << 7
	// IgnoreRomanNumerals ignores roman numerals
	IgnoreRomanNumerals = 2 << 8
)

var masks = map[string]Bitmask{
	"IGNORE_UPPERCASE":      IgnoreUppercase,
	"IGNORE_DIGITS":         IgnoreDigits,
	"IGNORE_URLS":           IgnoreUrls,
	"FIND_REPEAT_WORDS":     FindRepeatWords,
	"IGNORE_LATIN":          IgnoreLatin,
	"NO_SUGGEST":            NoSuggest,
	"FLAG_LATIN":            FlagLatin,
	"BY_WORDS":              ByWords,
	"IGNORE_CAPITALIZATION": IgnoreCapitalization,
	"IGNORE_ROMAN_NUMERALS": IgnoreRomanNumerals,
}

// Bitmask is used for bitmasks
type Bitmask int

func (b *Bitmask) String() string {
	return fmt.Sprint(*b)
}

// Set implements interface of flag.Value (https://golang.org/pkg/flag/#Value)
func (b *Bitmask) Set(value string) error {
	for _, key := range strings.Split(value, ",") {
		*b += masks[key]
	}

	return nil
}
