package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
	if (layerPrepTime == 0) {
		layerPrepTime = 2
	}

	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64){	
	for _, layer := range layers {
		if (layer == "noodles") {
			noodles += 50
		} else if ( layer == "sauce") {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList) - 1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var adjQuantities []float64
	adjPortion := float64(portions) / 2

	for _, quantity := range quantities {
		adjQuantities = append(adjQuantities, quantity * float64(adjPortion))
	}

	return adjQuantities
}
