package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"
)

//go:embed nsf2020.txt
var wordlist string

func parseWords() []string {
	return strings.Split(wordlist, "\n")
}

func canBuild(word string, letters string) bool {
	for _, c := range strings.Split(letters, "") {
		if strings.Count(word, c) > strings.Count(letters, c) {
			return false
		}
	}

	for _, c := range strings.Split(word, "") {
		if strings.Count(word, c) > strings.Count(letters, c) {
			return false
		}
	}
	return true
}

func findAllBuildable(letters string, legalWords []string) []string {
	var out []string
	for _, word := range legalWords {
		if canBuild(word, letters) {
			out = append(out, word)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return len(out[i]) < len(out[j])
	})

	return out
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: goscrabble letters")
		return
	}
	letters := os.Args[1]

	legalWords := parseWords()

	fmt.Println("Checking:", letters)
	matches := findAllBuildable(letters, legalWords)

	for _, match := range matches {
		fmt.Println(match)
	}

}
