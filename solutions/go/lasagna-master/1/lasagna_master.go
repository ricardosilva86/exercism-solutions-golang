package lasagna

// PreparationTime TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	n, s := 0, 0
	for _, v := range layers {
		if v == "noodles" {
			n += 1
		}

		if v == "sauce" {
			s += 1
		}
	}
	return n * 50, float64(s) * 0.2
}

// AddSecretIngredient TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
	return myList
}

// ScaleRecipe TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (newRecipe []float64) {
	for _, v := range quantities {
		newRecipe = append(newRecipe, (v/2)*float64(portions))
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
