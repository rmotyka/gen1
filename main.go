package main

import (
	"fmt"
	"gen1/selection"
	"gen1/cities"
	"gen1/initalPopulation"
	"gen1/mutation"
	"gen1/crossing"
	"sort"
	"gen1/structs"
	"gen1/graphics"
)


const populationLength = 100
const numberOfCities = 10
const maxCoordinate = 100

const enableMutation = false

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("gen2")

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

	// estimation
	for _, item := range population {
		item.Estimate(cityList, distances)
	}

	for _, item := range population	{
		fmt.Println(*item)
	}

	graphics.SaveCityImage(maxCoordinate, cityList)

	for cycleNumber := 0; cycleNumber < 10 ; cycleNumber++ {
		fmt.Println("Number of cycle ", cycleNumber)

		// selection
		newPopulation := selection.Select(population, populationLength)

		sort.Sort(structs.ByLenght(newPopulation))
		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("new population")
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		// crossing
		crossing.Crossing(newPopulation, numberOfCities)

		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("after crossing")
		sort.Sort(structs.ByLenght(newPopulation))
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		if enableMutation {
			// mutation
			mutation.Mutate(newPopulation, numberOfCities)

			for _, item := range newPopulation {
				item.Estimate(cityList, distances)
			}

			fmt.Println("after mutation")
			for _, item := range newPopulation {
				fmt.Println(*item)
			}
		}

		// repeat
		population = newPopulation

		// make graph
		citiesIds := structs.SelectCitiesIds(cityList, population[0].CitySelectionOrder)
		cities := make([]structs.City, len(citiesIds))
		for i, cityId := range citiesIds  {
			cities[i] = cityList[cityId]
		}

		graphics.SaveCityImage(maxCoordinate, cities)
	}
}
