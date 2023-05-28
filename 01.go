package main

import "fmt"

func canBuyEnoughFood(numberDogs, numberCats int, stock map[string]int) bool {
	totalNumberAnimals := numberDogs + numberCats
	amountFoodNeeded := stock["dog"] + stock["cat"] + stock["universal"]

	if amountFoodNeeded < totalNumberAnimals {
		return false
	}

	return true
}

func main() {
	numberDogs := 3
	numberCats := 3
	stock := map[string]int{
		"dog":       2,
		"cat":       2,
		"universal": 1,
	}

	possible := canBuyEnoughFood(numberDogs, numberCats, stock)
	fmt.Print(possible)
}
