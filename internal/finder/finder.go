package finder

import (
	"fmt"
	prmt "github.com/gitchander/permutation"
	"log"
	"strings"
)

var dictionary Dictionary

func init() {
	dictionary = Dictionary{}
	if err := dictionary.Load(); err != nil {
		log.Fatalln(err.Error())
	}
}

func FindWordsByLetters(letters string) ([]string, error) {
	if len(letters) < 1 {
		return nil, fmt.Errorf("the list of letter should contains at least 1 item")
	}
	letterList := strings.Split(letters, "")
	permutations := prmt.New(prmt.StringSlice(letterList))

	var results []string

	for permutations.Next() {
		results = append(results, dictionary.Search(strings.Join(letterList, ""))[:]...)
	}

	return unique(results), nil
}

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
