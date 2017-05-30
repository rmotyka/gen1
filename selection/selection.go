package selection

import (
	"gen1/structs"
	"math/rand"
	"fmt"
)

func Select(population []*structs.Route, populationLength int) []*structs.Route {
	newPopulation := make([]*structs.Route, len(population))
	// get the best chromosome and be sure to pass it further
	bestRoute := SelectBestItem(population)
	newPopulation[0] = &bestRoute

	fmt.Println("THE BEST IS ", bestRoute)

	for i := 1; i < populationLength; i++ {
		item := selectFromPopulation(population)
		newPopulation[i] = &item
	}

	return newPopulation
}

func SelectBestItem(population []*structs.Route ) structs.Route {
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
