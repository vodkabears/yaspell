package checker

import (
	"fmt"

	"github.com/vodkabears/yaspell/reader"
)

// Error stores misspells of a chunk
type Error struct {
	Chunk     *reader.Chunk
	Misspells []Misspell
}

func (e Error) Error() string {
	var str string
	file := e.Chunk.File
	lastIndex := len(e.Misspells) - 1
	for i, m := range e.Misspells {
		str += fmt.Sprintf("%s:%d:%d: %s", file, m.Row+1, m.Col+1, m.Word)
		if len(m.Suggestions) > 0 {
			str += fmt.Sprintf(" %v", m.Suggestions)
		}
		if i != lastIndex {
			str += "\n"
		}
	}

	return str
}
