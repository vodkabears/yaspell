package config_test

import (
	"strconv"
	"testing"

	"github.com/vodkabears/yaspell/config"
)

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

func TestBitmaskString(t *testing.T) {
	sum := config.IgnoreUppercase | config.IgnoreDigits
	b := config.Bitmask(sum)

	if b.String() != strconv.Itoa(sum) {
		t.Error("Expected to get sum of bitmasks")
	}
}
