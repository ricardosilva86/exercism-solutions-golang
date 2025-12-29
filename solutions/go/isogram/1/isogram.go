package isogram

import (
	"strings"
)

// import "strings"

func IsIsogram(word string) bool {
	// brute force solution using a map
	// then counting if any letter has more than 1 occurrence
	letters := make(map[string]int)
	for _, letter := range word {
		lowerCaseLetter := strings.ToLower(string(letter))
		letters[lowerCaseLetter]++
	}
	for l, count := range letters {
		if l != "-" && l != " " {
			if count > 1 {
				return false
			}
		}
	}
	return true

}
