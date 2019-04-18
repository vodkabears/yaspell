package checker_test

import (
	"github.com/vodkabears/yaspell/checker"
	"testing"
)

func TestRemoveIgnoredLines(t *testing.T) {
	str := "string_1\nstring_2\n// yaspell-disable-next-line\nstring_3\nstring_4"
	expected := "string_1\nstring_2\nstring_4"

	result := checker.RemoveIgnoredLines(str)

	if result != expected {
		t.Errorf("Expected to remove ignore lines.\nResult:\n%s \n\nExpected:\n%s", result, expected)
	}
}
