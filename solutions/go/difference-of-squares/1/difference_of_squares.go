package diffsquares

import (
    "fmt"
    "math"
    )

func SquareOfSum(n int) int {
    sum := 0
	for i:=0; i <= n; i++ {
        sum += i
        fmt.Println("sum:", sum)
    }
    power := math.Pow(float64(sum), float64(2))
    return int(power)
}

func SumOfSquares(n int) int {
	sq := 0.0
    for i:=0; i <= n; i++ {
        power := math.Pow(float64(i), float64(2))
        sq += power
    }

    return int(sq)
}

func Difference(n int) int {
	sqSum := SquareOfSum(n)
    sumSq := SumOfSquares(n)
    return sqSum - sumSq
}
