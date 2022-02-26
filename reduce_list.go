package main

import (
	"strings"
)

func ReduceList(word, matches, nonmatches string, perfMatches map[int]string, word_list []string) []string {
	word_list = removeWord(word, word_list)
	if len(nonmatches) > 0 {
		word_list = removeMatches(nonmatches, word_list)
	}
	for pos, ch := range perfMatches {
		word_list = keepOnlyPerfMatches(pos, ch, word_list)
	}

	if len(matches) > 0 {
		word_list = keepOnlyMatches(matches, word_list)
	}
	return word_list
}

func removeWord(word string, word_list []string) []string {
	for i, v := range word_list {
		if v == word {
			return append(word_list[:i], word_list[i+1:]...)
		}
	}
	return word_list
}

func removeMatches(nonmatches string, word_list []string) []string {
	new_list := make([]string, 0)

	for _, word := range word_list {
		match := false
		for i := 0; i < len(nonmatches) && !match; i++ {
			if strings.Contains(word, nonmatches[i:i+1]) {
				match = true
			}
		}
		if !match {
			new_list = append(new_list, word)
		}
	}
	return new_list
}

func keepOnlyPerfMatches(pos int, ch string, word_list []string) []string {
	new_list := make([]string, 0)
	for _, word := range word_list {
		if word[pos-1] == ch[0] {
			new_list = append(new_list, word)
		}
	}
	return new_list
}

func keepOnlyMatches(matches string, word_list []string) []string {
	new_list := make([]string, 0)
	for _, word := range word_list {
		matched := true
		for i := 0; i < len(matches) && matched; i++ {
			if !strings.Contains(word, matches[i:i+1]) {
				matched = false
			}
		}
		if matched {
			new_list = append(new_list, word)
		}
	}
	return new_list
}
