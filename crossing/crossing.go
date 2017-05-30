package crossing

import (
	"gen1/structs"
	"math/rand"
	"sort"
	"fmt"
)

func Crossing(population []*structs.Route, numberOfCities int) {
	sort.Sort(structs.ByLenght(population))

	for i := 0; i < 10 ; i++ {
		parentAIndex := rand.Intn(len(population) / 2)
		parentBIndex := rand.Intn(len(population))

		fmt.Println("crossing indexeses ", parentAIndex, parentBIndex)

		if parentAIndex == 0 || parentBIndex == 0 {
			continue
		}

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

