// Name: Sean Duncan
// NetID: sdduncan
// Description: The solution to Question 1 of Assignment 1-1. This file finds the most frequent words
// above a certain character threshold within a file
package cos418_hw1_1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuation and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// Compile Regex
	re, err := regexp.Compile(`[0-9a-zA-Z]+`)
	checkError(err)

	// Map words to wordcount
	wordMap := make(map[string]WordCount)

	// Open file check for err
	file, err := os.Open(path)
	defer file.Close()
	checkError(err)

	// wrap file in bufio scanner, facilitating reading file in line by line
	// and ensuring full contents have been read
	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() {
		line := fileReader.Text()

		line = strings.ToLower(line)
		fields := strings.Fields(line)
		for _, rawWord := range fields {
			word := strings.Join(re.FindAllString(rawWord, len(rawWord)), "")
			if len(word) < charThreshold {
				continue
			}

			if _, containsWord := wordMap[word]; !containsWord {
				wordMap[word] = WordCount{
					Word:word,
					Count:1,
				}
			} else {
				wc := wordMap[word]
				wc.Count++
				wordMap[word] = wc
			}
		}
	}

	var wordCounts []WordCount
	for _, key := range wordMap {
		wordCounts = append(wordCounts, key)
	}

	// Sort and return numWords number of words
	sortWordCounts(wordCounts)
	return wordCounts[0:numWords]
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
