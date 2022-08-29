package main

import (
	"fmt"
	"github.com/henri9813/findwordfromletters/internal/finder"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s letters\n", os.Args[0])
		log.Println(os.Args)
		os.Exit(2)
	}

	results, err := finder.FindWordsByLetters(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Words found:")
	for _, result := range results {
		log.Println(result)
	}
}
