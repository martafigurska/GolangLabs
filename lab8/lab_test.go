package main

import (
	"testing"
	"reflect"
)

func TestSplitWords(t *testing.T) {
	result, _ := SplitWords("Mam na imie Marta")
	expectedResult := []string{"Mam", "na", "imie", "Marta"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result was %v, expected result is %v", result, expectedResult)
	}
	_, err := SplitWords("test", ",", ";")
	if err == nil {
		t.Error("Should return error for too many params")
	}
}

func TestSplitSortWords(t *testing.T) {
	result, _ := SplitSortWords("d a c b")
	expectedResult := []string{"a", "b", "c", "d"}
	if len(result) != len(expectedResult) {
        t.Fatalf("Expected length %d, but got %d", len(expectedResult), len(result))
    }
	for i := range expectedResult {
        if result[i] != expectedResult[i] {
			t.Errorf("At index %d: expectedResult %s, got %s", i, expectedResult[i], result[i])
        }
	}
	_, err := SplitSortWords("test", ",", ";")
	if err == nil {
		t.Error("Should return error for too many params")
	}
}

func TestCountWords(t *testing.T) {
	result, _ := CountWords("Mam na imie Marta")
	if result != 4 {
		t.Errorf("Expected 4 words, got %d", result)
	}
}

func TestCountCharacters(t *testing.T) {
	result, _ := CountCharacters("Ala ma kota")
	if result != 9 {
		t.Errorf("Expected 9 characters, got %d", result)
	}
}

func TestWordFrequency(t *testing.T) {
	result, _ := WordFrequency("Ala ala ma ma ma kota")
	expectedResult := map[string]int{"ala":2, "ma":3, "kota":1}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result was %v, expected result is %v", result, expectedResult)
	}
}

func TestSortPalindroms(t *testing.T) {
	// 1 case
	result, _ := SortPalindroms("ala mam kota")
	expectedResult := []string{"ala", "mam"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result was %v, expected result is %v", result, expectedResult)
	}
	//2 case
	result, _ = SortPalindroms("tat mat aba")
	expectedResult = []string{"aba", "tat"}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Result was %v, expected result is %v", result, expectedResult)
	}
}

