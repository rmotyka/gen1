package crossing

import (
	"gen1/structs"
	"math/rand"
	"sort"
)

func Crossing(population []*structs.Route, numberOfCities int) {
	sort.Sort(structs.ByLenght(population))

	populationLen := len(population)
	maxIndexForA := populationLen / 4
	maxIndexForB := populationLen / 2

	i := 0
	for i < 10 {
		parentAIndex := rand.Intn(maxIndexForA)
		parentBIndex := rand.Intn(maxIndexForB - maxIndexForA) + maxIndexForA

		if parentAIndex == 0 || parentBIndex == 0 {
			continue
		}

		crossingPoint := rand.Intn(numberOfCities)

		routeA := population[parentAIndex]
		routeB := population[parentBIndex]

		//if routeA.Equals(routeB) {
		//	continue
		//}

		//childOrder1 := append(routeA.CitySelectionOrder[:crossingPoint], routeB.CitySelectionOrder[crossingPoint:]...)
		childOrder2 := append(routeB.CitySelectionOrder[:crossingPoint], routeA.CitySelectionOrder[crossingPoint:]...)

		// replace parents
		//routeA.CitySelectionOrder = childOrder1
		//routeA.Length = 0

		routeB.CitySelectionOrder = childOrder2
		routeB.Length = 0
		i++
	}
}

