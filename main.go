package main

import (
	"fmt"
	"math/rand"
)

type City struct {
	X int
	Y int
}

type Route struct {
	CitySelectionOrder []int
}

func (r *Route) Estimate(cityList map[int]City) {
	// select cities
	//cityNumberList := make([]int, numberOfCities)
	//for i := 0; i < numberOfCities; i++ {
	//
	//}
}

const populationLength = 10
const numberOfCities = 10
const maxCoordinate = 100

func main() {
	fmt.Println("gen2")

	cityList := getCityList()
	fmt.Println(cityList)

	population:=generateInitalPopulation()
	fmt.Println(population)

	for _, item:=range population  {
		item.Estimate(cityList)
	}

}

func generateInitalSelectionOrder() []int {
	citySelectionOrder:=make([]int, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		citySelectionOrder[i] = rand.Intn(numberOfCities-i)
	}

	return citySelectionOrder
}

func getCityList() (map[int]City) {
	cityList := make(map[int]City, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		cityList[i] = City{rand.Intn(maxCoordinate),rand.Intn(maxCoordinate)}
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
