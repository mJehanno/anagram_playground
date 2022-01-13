package main

import "testing"

type TestCase struct {
	Val1, Val2 string
	Expected   bool
}

var cases = []TestCase{
	{
		"palier",
		"pareil",
		true,
	},
	{
		"pareil",
		"joyeux",
		false,
	},
}

func TestAnagram(t *testing.T) {
	for _, c := range cases {
		result := Anagram(c.Val1, c.Val2)
		if result != c.Expected {
			t.Errorf("anagram test failed : expected %v, got %v with %s and %s", c.Expected, result, c.Val1, c.Val2)
		}
	}
}
