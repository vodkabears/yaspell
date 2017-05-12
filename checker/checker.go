package checker

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/VodkaBears/yaspell/config"
	"github.com/VodkaBears/yaspell/reader"
)

const serviceURL = "http://speller.yandex.net/services/spellservice.json/checkText"

// Misspell contains detailed information about a misspell
type Misspell struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

func checkWordByDictionary(word string, dictWords []*regexp.Regexp) bool {
	for _, dictWord := range dictWords {
		if dictWord.MatchString(word) {
			return true
		}
	}

	return false
}

// Check checks misspells in a chunk of text
func Check(chunk *reader.Chunk, cfg *config.Config) error {
	resp, err := http.PostForm(serviceURL, url.Values{
		"text":    {chunk.Text},
		"lang":    {cfg.Lang},
		"format":  {cfg.Format},
		"options": {strconv.Itoa(int(cfg.Options))},
	})
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
	}

	var misspells []Misspell
	if err = json.Unmarshal(body, &misspells); err != nil {
		return err
	}

	var dictWords = cfg.Dictionary.Words
	if len(dictWords) == 0 && len(misspells) > 0 {
		return Error{chunk, misspells}
	}

	var filteredMisspells []Misspell
	for _, misspell := range misspells {
		if !checkWordByDictionary(misspell.Word, dictWords) {
			filteredMisspells = append(filteredMisspells, misspell)
		}
	}

	if len(filteredMisspells) > 0 {
		return Error{chunk, filteredMisspells}
	}

	return nil
}
