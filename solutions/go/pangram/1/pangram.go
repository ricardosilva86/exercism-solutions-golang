package pangram

import (
	"maps"
	"slices"
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	if len(input) == 0 {
		return false
	}
	alphabet := make(map[string]int)

	for i := 0; i < len(input); i++ {
		letter := strings.ToLower(string(input[i]))
		if unicode.IsLetter(rune(letter[0])) {
			alphabet[letter]++
		}

	}

	keys := slices.Sorted(maps.Keys(alphabet))
	if len(keys) != 26 {
		return false
	}

	return true
}
