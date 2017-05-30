package selection

import (
	"gen1/structs"
	"math/rand"
	"fmt"
	"sort"
)

func Select(population []*structs.Route, populationLength int) []*structs.Route {
	sort.Sort(structs.ByLenght(population))

	newPopulation := make([]*structs.Route, len(population))
	// get the best chromosome and be sure to pass it further
	bestRoute := *population[0];
	newPopulation[0] = &bestRoute

	fmt.Println("THE BEST IS ", bestRoute)

	for i := 1; i < populationLength; i++ {
		item := selectFromPopulation(population)
		newPopulation[i] = &item
	}

	return newPopulation
}

func selectFromPopulation(population []*structs.Route) structs.Route {
	const tournamentSize = 5
	var bestRoute *structs.Route
	for i := 0; i < tournamentSize; i++ {
		itemIndex := rand.Intn(len(population))
		selectedItem := population[itemIndex]
		if bestRoute == nil || selectedItem.Length < bestRoute.Length {
			bestRoute = selectedItem
		}
	}

	return *bestRoute
}
