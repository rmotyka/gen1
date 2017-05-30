package selection

import (
	"gen1/structs"
	"math/rand"
)

func Select(population []*structs.Route, populationLength int) []*structs.Route {
	newPopulation := make([]*structs.Route, len(population))
	for i := 0; i < populationLength; i++ {
		item := selectFromPopulation(population)
		newPopulation[i] = item
	}

	return newPopulation
}

func selectFromPopulation(population []*structs.Route) *structs.Route {
	const tourneySize = 10
	var bestRoute *structs.Route
	for i := 0; i < tourneySize; i++ {
		itemIndex := rand.Intn(len(population))
		selectedItem := population[itemIndex]
		if bestRoute == nil || selectedItem.Length < bestRoute.Length {
			bestRoute = selectedItem
		}
	}

	return bestRoute
}
