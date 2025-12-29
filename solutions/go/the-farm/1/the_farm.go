package thefarm

import "fmt"

type InvalidCowsError struct {
	NumberOfCows int
	Message string
}

func (c InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", c.NumberOfCows, c.Message)
}

// TODO: define the 'DivideFood' function
func DivideFood(foodCalc FodderCalculator, cows int) (float64, error) {
	
	if cows == 0 {
		return 0.0, fmt.Errorf("something went wrong")
	}
	fodder, err := foodCalc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}

	fattingFactor, err := foodCalc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}

	return fodder * fattingFactor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(foodCalc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, fmt.Errorf("invalid number of cows")
	}

	return DivideFood(foodCalc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			NumberOfCows: cows,
			Message: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			NumberOfCows: 0,
			Message: "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
