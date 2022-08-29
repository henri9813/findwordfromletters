package finder

import (
	"encoding/json"
	"os"
)

// Dictionary is a dictionary representation.
type Dictionary struct {
	Words []string
}

// FileWord is the representation of a line of the dictionary
type FileWord struct {
	Original   string `json:"original"`
	Normalized string `json:"normalized"`
}

func (dictionary *Dictionary) Load() error {
	file, _ := os.ReadFile("../../data/french.json")

	var words []FileWord

	if err := json.Unmarshal([]byte(file), &words); err != nil {
		return err
	}

	for _, word := range words {
		dictionary.Words = append(dictionary.Words, word.Original)
		dictionary.Words = append(dictionary.Words, word.Normalized)
	}

	return nil
}

func (dictionary Dictionary) Search(search string) []string {
	var results []string
	for _, word := range dictionary.Words {
		if search == word {
			results = append(results, word)
		}
	}

	return results
}
