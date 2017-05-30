package crossing

import (
	"gen1/structs"
	"math/rand"
)

func Crossing(population []*structs.Route, numberOfCities int) {
	//bestRoute := selection.SelectBestItem(population)

	numberOfCrossingItems := int(len(population)*10/100) * 2
	itemsToCross := make([]*structs.Route, numberOfCrossingItems)

	crossingIndexes := crossingIndexes(numberOfCrossingItems, len(population)-1)

	for i, index := range crossingIndexes {
		item := population[index]
		itemsToCross[i] = item
	}

	for i := 0; i < numberOfCrossingItems/2; i++ {
		routeA := population[crossingIndexes[i]]
		routeB := population[crossingIndexes[i+numberOfCrossingItems/2]]
		crossingPoint := rand.Intn(numberOfCities)

		childOrder1 := append(routeA.CitySelectionOrder[:crossingPoint], routeB.CitySelectionOrder[crossingPoint:]...)
		childOrder2 := append(routeB.CitySelectionOrder[:crossingPoint], routeA.CitySelectionOrder[crossingPoint:]...)

		routeA.CitySelectionOrder = childOrder1
		routeA.Length = 0

		routeB.CitySelectionOrder = childOrder2
		routeB.Length = 0
	}
}

func crossingIndexes(numberOfCrossingItems int, maxIndex int) []int {
	indexes := make([]int, numberOfCrossingItems)
	for i := 0; i<numberOfCrossingItems; i++ {
		genOk := false
		for !genOk {
			genOk = true
			newIndex := rand.Intn(maxIndex)
			for j := 0; j < i; j++ {
				if newIndex == indexes[j] {
					genOk = false
					break
				}
			}

			if genOk {
				indexes[i] = newIndex
			}
		}

	}

	return indexes
}