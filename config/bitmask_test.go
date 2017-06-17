package config_test

import (
	"strconv"
	"testing"

	"github.com/vodkabears/yaspell/config"
)

const totalBitmaskSum = 3007

func TestBitmaskSet(t *testing.T) {
	var b config.Bitmask

	err := b.Set("IGNORE_UPPERCASE,IGNORE_DIGITS,IGNORE_URLS")
	if err != nil {
		t.Log("Unexpected error:", err)
	}

	if b != config.IgnoreUppercase|config.IgnoreDigits|config.IgnoreUrls {
		t.Error("Expected to be equal to sum of bitmasks")
	}
}

func TestBitmaskTotalSum(t *testing.T) {
	b := config.Bitmask(config.IgnoreUppercase |
		config.IgnoreDigits |
		config.IgnoreUrls |
		config.FindRepeatWords |
		config.IgnoreLatin |
		config.NoSuggest |
		config.FlagLatin |
		config.ByWords |
		config.IgnoreCapitalization |
		config.IgnoreRomanNumerals)

	if b != totalBitmaskSum {
		t.Errorf("Expected to be equal %d", totalBitmaskSum)
	}
}

func TestBitmaskString(t *testing.T) {
	sum := config.IgnoreUppercase | config.IgnoreDigits
	b := config.Bitmask(sum)

	if b.String() != strconv.Itoa(sum) {
		t.Error("Expected to get sum of bitmasks")
	}
}
