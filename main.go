package main

import (
	"bufio"
	"fmt"
	"gordle/gordle"
	"os"
)

const maxAttempts = 6
const corpusFile = "./corpus/engish.txt"

func main() {
	corpus, err := gordle.ReadCorpus(corpusFile)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	// create a new game
	g, err := gordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "uanble to start game: %s", err)
		return
	}

	g.Play()
}
