package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	c := 0
	sum := n

	if n == 0 {
		return 0, errors.New("zero is an error")
	}

	if n < 0 {
		return 0, errors.New("negative value is an error")
	}

	for {
		if sum == 1 {
			return c, nil
		}

		if sum%2 == 0 {
			sum = sum / 2
		} else {
			sum = 3*sum + 1
		}
		c++
	}
}
