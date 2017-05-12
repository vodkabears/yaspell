package reader_test

import (
	"io/ioutil"
	"testing"

	"github.com/VodkaBears/yaspell/reader"
)

func TestReaderRead(t *testing.T) {
	ch := make(chan reader.Chunk)
	go reader.Read(ch, []string{"testdata/text1.txt", "testdata/text2.txt"})

	for chunk := range ch {
		text, _ := ioutil.ReadFile(chunk.File)
		if string(text) != chunk.Text {
			t.Error("Expected to correctly read files")
			break
		}
	}
}
