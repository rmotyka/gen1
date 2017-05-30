package mutation

import (
	"gen1/structs"
	"math/rand"
	"sort"
)

func Mutate(population []*structs.Route, numberOfCities int) {
	sort.Sort(structs.ByLenght(population))

	numberOfMutations := int(len(population) * 10/100);
	for i :=0; i<numberOfMutations; i++ {
		itemIndex := rand.Intn(len(population))
		if itemIndex == 0 {
			continue
		}

		item := population[itemIndex]

		cityIndex := rand.Intn(numberOfCities)
		newCityIndex := rand.Intn(numberOfCities - cityIndex)
		item.CitySelectionOrder[cityIndex] = newCityIndex
		item.Length = 0
	}
}
