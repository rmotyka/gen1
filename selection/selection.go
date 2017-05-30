package selection

import (
	"gen1/structs"
	"math/rand"
)

func Select(population []*structs.Route, populationLength int) []*structs.Route {
	newPopulation := make([]*structs.Route, len(population))
	// get the best chromosome and be sure to pass it further
	bestRoute := selectBestItem(population)
	newPopulation[0] = &bestRoute

	for i := 1; i < populationLength; i++ {
		item := selectFromPopulation(population)
		newPopulation[i] = &item
	}

	return newPopulation
}

func selectBestItem(population []*structs.Route ) structs.Route {
	bestItem := population[0]
	for i := 1; i < len(population); i++ {
		if population[i].Length < bestItem.Length {
			bestItem = population[i]
		}
	}

	return *bestItem
}

func selectFromPopulation(population []*structs.Route) structs.Route {
	const tourneySize = 5
	var bestRoute *structs.Route
	for i := 0; i < tourneySize; i++ {
		itemIndex := rand.Intn(len(population))
		selectedItem := population[itemIndex]
		if bestRoute == nil || selectedItem.Length < bestRoute.Length {
			bestRoute = selectedItem
		}
	}

	return *bestRoute
}
