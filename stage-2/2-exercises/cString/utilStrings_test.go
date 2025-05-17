package cString_test

import (
	"exercise/cString"
	"testing"
)

func TestCountVowels(t *testing.T) {
	var tests = []struct {
		name   string
		vowels string
		want   int
	}{
		{"zero vowels", "", 0},
		{"one vowel", "a", 1},
		{"two vowels", "ae", 2},
		{"three vowels", "aei", 3},
		{"four vowels", "aeio", 4},
		{"five vowels", "aeiou", 5},
	}

	for _, tt := range tests {
		testName := tt.name
		t.Run(testName, func(t *testing.T) {
			got := cString.CountVowels(tt.vowels)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCapitalizeAndRemoveSpaces(t *testing.T) {
	var tests = []struct {
		name   string
		phrase string
		want   string
	}{
		{"all lowered", "hello world now", "Hello World Now"},
		{"starting with number", "1hello world now", "1hello World Now"},
		{"mixed cases", "hellO worLd now", "Hello World Now"},
		{"more than one space", "hello      worLd now", "Hello World Now"},
		{"cleaning tabs", "hello\t\tworLd now", "Hello World Now"},
		{"cleaning new lines", "hello\n\tworLd now", "Hello World Now"},
	}

	for _, tt := range tests {
		testName := tt.name
		t.Run(testName, func(t *testing.T) {
			got := cString.CapitalizeAndRemoveSpaces(tt.phrase)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
