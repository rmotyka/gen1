package crossing

import (
	"gen1/structs"
	"math/rand"
)

func Crossing(population []*structs.Route, numberOfCities int) {
	numberOfCrossing := int(len(population) * 10/100);
	for i :=0; i<numberOfCrossing; i++ {
		parentAIndex := rand.Intn(len(population))
		parentBIndex := rand.Intn(len(population))
		crossingPoint := rand.Intn(numberOfCities)

		routeA := population[parentAIndex]
		routeB := population[parentBIndex]

		childOrder1 := append(routeA.CitySelectionOrder[:crossingPoint], routeB.CitySelectionOrder[crossingPoint:]...)
		childOrder2 := append(routeB.CitySelectionOrder[:crossingPoint], routeA.CitySelectionOrder[crossingPoint:]...)

		// replace parents
		routeA.CitySelectionOrder = childOrder1
		routeA.Length = 0

		routeB.CitySelectionOrder = childOrder2
		routeB.Length = 0
	}
}