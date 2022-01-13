package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	r := Anagrams(os.Args[1:])
	for _, v := range r {
		fmt.Println(v)
	}
}

func Anagrams(input []string) map[string][]string {
	m := map[string]string{}
	r := map[string][]string{}
	for _, refWord := range input {
		m[refWord] = FindCount(refWord)
	}

	for k, v := range m {
		r[v] = append(r[v], k)
	}

	return r
}

func FindCount(word string) string {
	return SortString(word)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
