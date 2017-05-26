package main

import (
	"fmt"
	"math/rand"
	"math"
)

type City struct {
	Id int
	X int
	Y int
}

type Route struct {
	CitySelectionOrder []int
}

func (r *Route) Estimate(cityList []City) {
	cityListCopy := make([]City, len(cityList))
	copy(cityListCopy, cityList)
	cities:=selectCities(cityListCopy, r.CitySelectionOrder)
	fmt.Println("City in Route")
	fmt.Println(cities)
}

func selectCities(cityList []City, citySelectionOrder []int) []City {
	outCityList := make([]City, len(citySelectionOrder))
	for i:=0; i<len(citySelectionOrder); i++ {
		indexOfCity := citySelectionOrder[i]
		city := cityList[indexOfCity]
		outCityList[i] = city

		// remove indexOfCity from cityList:
		cityList = append(cityList[:indexOfCity], cityList[indexOfCity+1:]...)
	}

	return outCityList
}


const populationLength = 10
const numberOfCities = 10
const maxCoordinate = 100

func main() {
	fmt.Println("gen2")

	cityList := getCityList()
	fmt.Println("Inital city list")
	fmt.Println(cityList)

	// calculate distances
	distances := calculateDistances(cityList)
	fmt.Println("Distances")
	fmt.Println(distances)

	population:=generateInitalPopulation()
	fmt.Println("Population")
	fmt.Println(population)

	for _, item:=range population  {
		item.Estimate(cityList)
	}

}

func calculateDistances(cityList []City) [][]float64 {
	distances := make([][]float64, len(cityList))
	for i, cityFrom := range cityList {
		distances[i] = make([]float64, len(cityList))
		for j, cityTo := range cityList {
			distance := float64(0)
			if cityFrom.Id != cityTo.Id {
				squareSum := math.Pow(float64(cityFrom.X-cityTo.X), 2)+math.Pow(float64(cityFrom.Y-cityTo.Y), 2)
				distance = math.Sqrt(squareSum)
			}

			distances[i][j] = distance
		}
	}

	return distances
}

func generateInitalSelectionOrder() []int {
	citySelectionOrder:=make([]int, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		citySelectionOrder[i] = rand.Intn(numberOfCities-i)
	}

	return citySelectionOrder
}

func getCityList() ([]City) {
	cityList := make([]City, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		cityList[i] = City{ i,rand.Intn(maxCoordinate),rand.Intn(maxCoordinate)}
	}

	return cityList;
}

func generateInitalPopulation() ([]Route) {
	population := make([]Route, populationLength)
	for i := 0; i < populationLength; i++ {
		citySelectionOrder :=generateInitalSelectionOrder()
		population[i] = Route{ citySelectionOrder }
	}

	return population
}
