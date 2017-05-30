package initalPopulation

import (
	"gen1/structs"
	"math/rand"
)

func generateInitialSelectionOrder(numberOfCities int) []int {
	citySelectionOrder:=make([]int, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		citySelectionOrder[i] = rand.Intn(numberOfCities-i)
	}

	return citySelectionOrder
}

func GenerateInitialPopulation(populationLength int, numberOfCities int) ([]*structs.Route) {
	population := make([]*structs.Route, populationLength)
	for i := 0; i < populationLength; i++ {
		citySelectionOrder := generateInitialSelectionOrder(numberOfCities)
		population[i] = &structs.Route{ citySelectionOrder, 0}
	}

	return population
}
