package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime == 0 {
		layerPrepTime = 2
	}
	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, ownrecipe []string) {
	ownrecipe[len(ownrecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledPortions []float64
	for _, quantity := range quantities {
		scaledPortions = append(scaledPortions, quantity*(float64(portions)/2))
	}
	return scaledPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
