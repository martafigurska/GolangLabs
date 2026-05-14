package main

import (
	"fmt"
	"strings"
	"slices"
	"errors"
)

func SplitWords(input string, separators ...string) ([]string, error) {
	if len(separators) > 1 {
		return nil, errors.New("Too many params")
	}
	var result []string
	if len(separators) == 0 {
		result = strings.Fields(input)
	} else {
		sep := separators[0]
		result = strings.Split(input, sep)
		var filtered_input []string
		for _, w := range result {
			if w != "" {
				filtered_input = append(filtered_input, w)
			}
		}
		result = filtered_input 
	}
	return result, nil
}

func SplitSortWords(input string, separators ...string) ([]string, error) {
	var result []string
	result, err := SplitWords(input, separators...)
	if err != nil {
		return nil, err
	}
	slices.Sort(result)

	return result, nil
}

func CountWords(input string) (int, error) {
	var result []string
	result, err := SplitWords(input)
	if err != nil {
		return 0, err
	}
	countWords := len(result)
	return countWords, err
}

func CountCharacters(input string) (int, error) {
	var result []string
	result, err := SplitWords(input)
	if err != nil {
		return 0, err
	}
	var countCharacters = 0
	for i := range result {
		countCharacters += len(result[i])
	}
	return countCharacters, nil
}

func WordFrequency(input string) (map[string]int, error) {
	var result []string
	input = strings.ToLower(input)
	result, err := SplitWords(input)
	if err != nil {
		return nil, err
	}
	freq := make(map[string]int)
	for _, word := range result {
		freq[word]++
	}
	return freq, nil
}

func SortPalindroms(input string) ([]string, error) {
	var result []string
	input = strings.ToLower(input)

	result, err := SplitWords(input)
	if err != nil {
		return nil, err
	}

	var palindroms []string
	for i := range result {
		for j := range len(result[i])/2 {
			wordLen := len(result[i])
			if result[i][j] != result[i][wordLen-j-1] {
				break
			}
		palindroms = append(palindroms, result[i])
		}
	}
	slices.Sort(palindroms)
	return palindroms, nil
}

func main() {
	text := "Ala\nma\nkota"
	fmt.Println(strings.Fields(text))
}