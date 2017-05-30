package mutation

import (
	"gen1/structs"
	"math/rand"
)

func Mutate(population []*structs.Route, numberOfCities int) {
	numberOfMutations := int(len(population) * 10/100);
	for i :=0; i<numberOfMutations; i++ {
		itemIndex := rand.Intn(len(population))
		item := population[itemIndex]

		cityIndex := rand.Intn(numberOfCities)
		newCityIndex := rand.Intn(numberOfCities - cityIndex)
		item.CitySelectionOrder[cityIndex] = newCityIndex
		item.Length = 0
	}
}
