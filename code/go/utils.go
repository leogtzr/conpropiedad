package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

func validateOptions(options ...string) bool {
	if len(options) == 0 {
		return false
	}
	r := true

	for _, opt := range options {
		if len(strings.TrimSpace(opt)) == 0 {
			r = false
			break
		}
	}
	return r
}

func loadInputFile(inputFile string, wordFormat *regexp.Regexp) ([]Word, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make([]Word, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordText := scanner.Text()
		if shouldIgnoreLine(wordText) {
			continue
		}
		if !isWordFormatValid(wordText, wordFormat) {
			continue
		}
		word := extractWordFromText(wordText)
		words = append(words, word)
	}

	return words, nil
}

func extractWordFromText(wordText string) Word {
	wordTokens := strings.Split(wordText, ";")
	word := Word{}
	word.word = wordTokens[0]
	word.meaning = wordTokens[1]
	tags := strings.Split(wordTokens[2], ",")
	word.tags = tags
	return word
}

func shouldIgnoreLine(line string) bool {
	return strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0
}

func isWordFormatValid(wordText string, rgx *regexp.Regexp) bool {
	return rgx.MatchString(wordText)
}

func equalTags(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalWords(a, b Word) bool {
	return a.word == b.word && a.meaning == b.meaning && equalTags(a.tags, b.tags)
}

func stringWithCharset(length int, charset string, seededRand *rand.Rand) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func filterByText(words *[]Word, text string) []Word {
	filteredWords := make([]Word, 0)

	for _, w := range *words {
		for _, tag := range w.tags {
			if strings.Contains(tag, text) {
				filteredWords = append(filteredWords, w)
				break
			}
		}
	}

	return filteredWords
}

func (w Word) String() string {
	return fmt.Sprintf("%s: %s\t\t(%s)", w.word, w.meaning, strings.Join(w.tags, ", "))
}
