package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"flag"
	"log"
	"os"
	"strconv"
)

func ReadData(r io.Reader, splitWords bool, prefix string, unique bool) []string {
	scanner := bufio.NewScanner(r)

	if splitWords {
		scanner.Split(bufio.ScanWords)
	}

	var res []string
	seen := make(map[string]bool)

	for scanner.Scan() {
		val := scanner.Text()
		if (prefix == "" || strings.HasPrefix(val, prefix)) && (!unique || !seen[val]) {
			res = append(res, val)
			seen[val] = true
		}
	}
	return res
}

func FunTrue[T any](data []T, fun func(T) bool) []T {
	var result []T
	for _, v := range data {
		if fun(v) {
			result = append(result, v)
		}
	}
	return result
}

func FunElements[T any](data []T, fun func(T) T) []T {
	var result []T
	for _, v := range data {
		result = append(result, fun(v))
	}
	return result
}

func MapToSlices[K comparable, V any](m map[K]V) ([]K, []V) {
	log.Println("Map to make separate slices :", m)
	var keys []K
	var values []V
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

type Pair[K comparable, V any] struct {
	Key K
	Value V
}

func MapToPairSlice[K comparable, V any](m map[K]V) []Pair[K, V] {
	log.Println("Map to make slice of pairs :", m)
	var pairs []Pair[K, V]
	for k, v := range m {
		pairs = append(pairs, Pair[K, V]{k, v})
	}
	return pairs
}

func SlicesToMap[K comparable, V any](k []K, v []V) map[K]V {
	log.Println("Slices to make a map :", k, v)
	slicesMap := make(map[K]V)
	lenMap := len(k)
	if lenMap > len(v) {
		lenMap = len(v)
	}

	for i := 0; i < lenMap; i++ {
		slicesMap[k[i]] = v[i] 
	}
	return slicesMap
}

func ConvertSlice[T int | float64](s []string) []T {
	var result []T
	for i := 0; i < len(s); i++ {
		var converted T
        switch any(converted).(type) {
        case int:
            v, _ := strconv.Atoi(s[i])
            converted = T(v)
        case float64:
            v, _ := strconv.ParseFloat(s[i], 64)
            converted = T(v)
        }
		result = append(result, converted)
	}
	return result
}

func main() {
	filePath := flag.String("file", "", "filepath")
	words := flag.Bool("w", false, "words")
	uniq := flag.Bool("u", false, "unique")
	pref := flag.String("p", "", "prefix")
	flag.Parse()
	
	if flag.NFlag() == 0 {
		log.Println("No flags")
	} else {
		flag.Visit(func(f *flag.Flag) {
		log.Printf("Used flag: -%s=%v\n", f.Name, f.Value)})
	}

	// fmt.Println("Flags:"*filePath, *words, *uniq, *pref)

	var input io.Reader = os.Stdin
	if *filePath != "" {
		file, err := os.Open(*filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		input = file
	}

	out := ReadData(input, *words, *pref, *uniq)
	fmt.Println("Result of flagging :", out)
	
	var funTrue = FunTrue(out, func(s string) bool {return len(s) == 2})
	log.Println("FunTrue result :", funTrue)

	var funElements = FunElements(out, func(s string) string {return "Element: " + s})
	log.Println("FunElements result :", funElements)

	labPoints := map[string]int{"lab1":6, "lab2":9,"lab3":7}

	labs, points := MapToSlices(labPoints)
	log.Println("Lab numbers :", labs)
	log.Println("Lab points :", points)

	labPointrPairs := MapToPairSlice(labPoints)
	log.Println("Lab + points :", labPointrPairs)

	numLetMap := SlicesToMap(labs, points)
	log.Println("Two slices into map results :", numLetMap)

	nums := []string{"1", "2"}
	stringToT := ConvertSlice[int](nums)
	log.Println("String to int/float :", stringToT)
}