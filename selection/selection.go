package selection

import (
	"gen1/structs"
	"math/rand"
)

func SelectFromPopulation(population []*structs.Route) structs.Route {
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
