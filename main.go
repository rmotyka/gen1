package main

import (
	"fmt"
	"gen1/cities"
	"gen1/initalPopulation"
	"gen1/graphics"
	"gen1/selection"
	"sort"
	"gen1/structs"
	"math/rand"
)


const populationLength = 100
const numberOfCities = 10
const maxCoordinate = 100

const enableMutation = false

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("gen3")

	cityList := cities.GetCityList(numberOfCities, maxCoordinate)
	fmt.Println("Inital city list")
	fmt.Println(cityList)

	// calculate distances
	distances := cities.CalculateDistances(cityList)
	fmt.Println("Distances")
	fmt.Println(distances)

	// population
	population:=initalPopulation.GenerateInitialPopulation(populationLength, numberOfCities)
	fmt.Println("Population")

	// sort population
	sort.Sort(structs.ByLenght(population))

	// estimation
	for _, item := range population {
		item.Estimate(cityList, distances)
	}
	// print population
	for _, item := range population	{
		fmt.Println(*item)
	}

	graphics.SaveCityImage("0", maxCoordinate, cityList)

	var bestRoute *structs.Route

	for cycleNumber := 0; cycleNumber < 1000 ; cycleNumber++ {
		fmt.Println("Number of cycle ", cycleNumber)

		parentA := selection.SelectFromPopulation(population)
		parentB := selection.SelectFromPopulation(population)
		if rand.Float32() < 0.8 {
			crossingPoint := rand.Intn(numberOfCities)

			childOrder1 := append(parentA.CitySelectionOrder[:crossingPoint], parentB.CitySelectionOrder[crossingPoint:]...)
			childOrder2 := append(parentB.CitySelectionOrder[:crossingPoint], parentA.CitySelectionOrder[crossingPoint:]...)

			// replace parents
			parentA.CitySelectionOrder = childOrder1
			parentA.Length = 0

			parentB.CitySelectionOrder = childOrder2
			parentB.Length = 0
		}

		if rand.Float32() < 0.1 {
			itemToMutate := selection.SelectFromPopulation(population)

			cityIndex := rand.Intn(numberOfCities)
			newCityIndex := rand.Intn(numberOfCities - cityIndex)
			itemToMutate.CitySelectionOrder[cityIndex] = newCityIndex
			itemToMutate.Length = 0
		}

		// sort population
		sort.Sort(structs.ByLenght(population))

		// estimation
		for _, item := range population {
			item.Estimate(cityList, distances)
		}
		// print population
		//for _, item := range population	{
		//	fmt.Println(*item)
		//}
		bestRoute = population[0]
		fmt.Println(*bestRoute)

		if bestRoute.Length < 300 {
			break
		}
	}


	citiesIds := structs.SelectCitiesIds(cityList, bestRoute.CitySelectionOrder)
	cities := make([]structs.City, len(citiesIds))
	for i, cityId := range citiesIds {
		cities[i] = cityList[cityId]
	}

	graphics.SaveCityImage("Final", maxCoordinate, cities)
}
