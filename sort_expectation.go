package main

import (
	"sort"
	"strings"
)

const (
	bull = 1
	hit  = 2
)

type WordExpectationPair struct {
	word string
	val  float64
}

func SortByExpectation(word_list []string) []string {
	wordExpectSlice := expectationOfWord(word_list[1], word_list)
	top10wordList := make([]string, 10)
	for i := 0; i < 10 && i < len(wordExpectSlice); i++ {
		top10wordList = append(top10wordList, wordExpectSlice[i].word)
	}

	return top10wordList
}

func expectationOfWord(word string, word_list []string) []WordExpectationPair {
	var weps []WordExpectationPair
	for _, word1 := range word_list {
		wordExpectation := make(map[int]int)
		for _, word2 := range word_list {
			if word1 != word2 {
				pattern := compare2words(word1, word2)
				if _, found := wordExpectation[pattern]; found {
					wordExpectation[pattern] = wordExpectation[pattern] + 1
				} else {
					wordExpectation[pattern] = 1
				}
			}
		}
		possiblePatterns := len(wordExpectation)
		eValue := 0.0
		for _, occurance := range wordExpectation {
			eValue += float64(occurance) / float64(possiblePatterns)
		}
		weps = append(weps, WordExpectationPair{word1, eValue})
	}
	sort.Slice(weps, func(i, j int) bool {
		return weps[i].val < weps[j].val
	})

	return weps
}

func compare2words(word1, word2 string) int {
	comparison := 0
	comparison += compareChar2Word(word1, word2, 0, 1, 10000)
	comparison += compareChar2Word(word1, word2, 1, 2, 1000)
	comparison += compareChar2Word(word1, word2, 2, 3, 100)
	comparison += compareChar2Word(word1, word2, 3, 4, 10)
	comparison += compareChar2Word(word1, word2, 4, 5, 1)
	return comparison
}

func compareChar2Word(word1, word2 string, from, to, digit int) int {
	comp := 0
	if word1[from:to] == word2[from:to] {
		comp += hit * digit
	} else if strings.Contains(word2, word1[from:to]) {
		comp += bull * digit
	}
	return comp
}
