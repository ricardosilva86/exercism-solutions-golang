package luhn

import (
    "strconv"
    "strings"
    "unicode"
    )

func Valid(id string) bool {
    numbers := strings.ReplaceAll(id, " ", "")
    if len(numbers) <= 1 { return false }

    for _, v := range numbers {
        if !unicode.IsDigit(v) { return false }
    }
    sum := 0
    doubleNext := false
    for i:=len(numbers)-1; i >= 0; i-- {
        double, _ := strconv.Atoi(string(numbers[i]))
        if doubleNext {
            double *= 2
            if double > 9 { double -= 9 }
        }
        sum += double
        doubleNext = !doubleNext
    }
	return sum%10==0
}
