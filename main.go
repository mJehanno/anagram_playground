package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println(Anagram(os.Args[1], os.Args[2]))
}

func Anagram(word1, word2 string) bool {
	count1, count2 := map[rune]int{}, map[rune]int{}
	for _, r := range word1 {
		count1[r]++
	}
	for _, r := range word2 {
		count2[r]++
	}
	return reflect.DeepEqual(count1, count2)
}
