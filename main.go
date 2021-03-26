package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const WordlistPath = "Path/To/Wordlist"

func readFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
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

	legalWords, err := readFile(WordlistPath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Checking:", letters)
	matches := findAllBuildable(letters, legalWords)

	for _, match := range matches {
		fmt.Println(match)
	}

}
