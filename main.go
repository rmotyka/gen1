package main

import (
	"fmt"
	//"time"
	"gen1/selection"
	"gen1/cities"
	"gen1/initalPopulation"
	"gen1/mutation"
	"gen1/crossing"
)


const populationLength = 100
const numberOfCities = 10
const maxCoordinate = 100

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

	for cycleNumber := 0; cycleNumber < 5 ; cycleNumber++ {
		fmt.Println("Number of cycle ", cycleNumber)

		// estimation
		for _, item := range population {
			item.Estimate(cityList, distances)
		}

		// selection
		newPopulation := selection.Select(population, populationLength)

		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("new population")
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		// crossing
		// TODO: verify if it works
		crossing.Crossing(newPopulation, numberOfCities)

		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("after crossing")
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		// mutation
		mutation.Mutate(newPopulation, numberOfCities)

		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("after mutation")
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		// repeat
		population = newPopulation
	}
}
