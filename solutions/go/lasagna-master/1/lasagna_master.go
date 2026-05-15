package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string , prep int) int {
    if prep == 0{
        return len(layers) * 2
    }
    return len(layers) * prep
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string,) ( noodles int, sauce float64){

    for _, layer := range layers{
        if layer == "sauce"{
            sauce += 0.2
        }else if layer == "noodles" {
            noodles += 50
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredientListFriend []string, ingredientListPersonal []string){
    lastIngredientIndex := len(ingredientListFriend) -1
    PersonalLastIngIndex := len(ingredientListPersonal) -1

    ingredientListPersonal[PersonalLastIngIndex] = ingredientListFriend[lastIngredientIndex]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities  []float64, portionNumber int) []float64{
    scaledQuantities := make([]float64, len(quantities))

    factor := float64(portionNumber) / 2.0
    
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * factor
    }
    
    return scaledQuantities
}
