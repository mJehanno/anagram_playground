package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	Values   []string
	Expected [][]string
}

var cases = []TestCase{
	{
		[]string{
			"atropine", "tarsiens", "palier", "nitrasse", "tapissera", "plaire", "pointera",
			"pareil", "parasites", "replia", "repassait", "pianoter", "perlai", "sentiras", "assirent",
			"passerait", "transies", "astreins",
		},
		[][]string{
			{"atropine", "pointera", "pianoter"},
			{"tarsiens", "nitrasse", "sentiras", "assirent", "transies", "astreins"},
			{"tapissera", "parasites", "repassait", "passerait"},
			{"palier", "plaire", "pareil", "replia", "perlai"},
		},
	},
}

func TestAnagram(t *testing.T) {
	for _, c := range cases {
		result := Anagrams(c.Values)
		if !reflect.DeepEqual(result, c.Expected) {
			t.Errorf("anagram test failed : expected %v, got %v", c.Expected, result)
		}
	}
}
