// Package word provides custom functions for working with strings and words.
package word

import "strings"

// UseCount counts the number of use for each word.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Counts the number of words in a string.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
