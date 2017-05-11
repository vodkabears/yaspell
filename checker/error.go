package checker

import (
	"fmt"

	"github.com/yaspell/reader"
)

// Error stores misspells of a chunk
type Error struct {
	Chunk     *reader.Chunk
	Misspells []Misspell
}

func (e Error) Error() string {
	file := e.Chunk.File
	lastIndex := len(e.Misspells) - 1
	var str string
	for i, m := range e.Misspells {
		str += fmt.Sprintf("%s:%d:%d: %s", file, m.Row, m.Col, m.Word)
		if len(m.Suggestions) > 0 {
			str += fmt.Sprintf(" %v", m.Suggestions)
		}
		if i != lastIndex {
			str += "\n"
		}
	}

	return str
}
