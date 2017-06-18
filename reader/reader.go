package reader

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

const maxBytes = 10000
const batchSize = 256

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

// Chunk is a chunk of text from file
type Chunk struct {
	File string
	Text string
}

func readFile(file string, ch chan Chunk, wg *sync.WaitGroup) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	b := bufio.NewReader(f)
	buf := make([]byte, maxBytes)
	for {
		_, err = b.Read(buf)
		if err == io.EOF {
			break
		}

		ch <- Chunk{file, string(bytes.Trim(buf, "\x00"))}
	}

	err = f.Close()
	if err != nil {
		log.Println(err)
	}

	wg.Done()
}

func readBatch(ch chan Chunk, files []string) {
	var wg sync.WaitGroup

	wg.Add(len(files))
	for _, file := range files {
		go readFile(file, ch, &wg)
	}

	wg.Wait()
}

// Read reads a batch of files chunk by chunk and passes content to the channel
func Read(ch chan Chunk, files []string) {
	for i, l := 0, len(files); i < l; i += batchSize {
		readBatch(ch, files[i:min(l, i+batchSize)])
	}

	close(ch)
}
