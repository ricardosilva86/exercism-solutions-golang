package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	numbers := strings.ReplaceAll(id, " ", "")
	if len(numbers) <= 1 {
		return false
	}

	for _, r := range numbers {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	sum := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(numbers[i]))
		if (len(numbers)-1-i)%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
