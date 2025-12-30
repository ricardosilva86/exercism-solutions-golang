package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	if len(input) == 0 {
		return false
	}

	seenLetters := [26]bool{} // One for each letter a-z

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			letterIndex := lowerChar - 'a'
			seenLetters[letterIndex] = true
		}
	}

	// Check if all letters a-z are present
	for _, seen := range seenLetters {
		if !seen {
			return false
		}
	}

	return true
}
