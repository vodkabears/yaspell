package checker

import (
	"math"
	"strings"
)

// IgnoreComment it's comment name
const IgnoreComment = "yaspell-disable-next-line"

// RemoveIgnoredLines removes comment and next line
func RemoveIgnoredLines(str string) string {
	ignoreLine := -math.MaxInt8
	temp := strings.Split(str, "\n")
	text := make([]string, 0)

	for idx, item := range temp {
		if strings.Contains(item, IgnoreComment) {
			ignoreLine = idx
		} else if ignoreLine+1 != idx {
			text = append(text, item)
			ignoreLine = -math.MaxInt8
		}
	}

	return strings.Join(text, "\n")
}
