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

const populationLength = 100
const numberOfCities = 3
const maxCoordinate = 100

func main() {
	fmt.Println("gen2")

	cityList := getCityList()
	fmt.Println(cityList)

	population:=generateInitalPopulation()
	fmt.Println(population)
}

func generateInitalSelectionOrder() []int {
	citySelectionOrder:=make([]int, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		citySelectionOrder[i] = rand.Intn(numberOfCities)
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
