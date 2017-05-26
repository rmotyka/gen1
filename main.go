package main

import (
	"fmt"
	"math/rand"
	"math"
	//"image"
	//"image/color"
	//"os"
	//"log"
	//"image/jpeg"
)

type City struct {
	Id int
	X int
	Y int
}

type Route struct {
	CitySelectionOrder []int
	Length float64
}

func (r *Route) Estimate(cityList []City, distances [][]float64) {
	cityListIds := make([]int, len(cityList))
	for i, city := range cityList  {
		cityListIds[i] = city.Id
	}

	citiesInRouteOrder :=selectCitiesIds(cityListIds, r.CitySelectionOrder)
	fmt.Println("City in Route")
	fmt.Println(citiesInRouteOrder)

	routeLenght := float64(0)
	previousCityId := -1
	for _, cityId := range citiesInRouteOrder  {
		if previousCityId != -1 {
			l := distances[cityId][previousCityId]
			routeLenght += l
		}

		previousCityId = cityId
	}

	r.Length = routeLenght
	fmt.Println("Route lenght ", routeLenght)
}

func selectCitiesIds(cityList []int, citySelectionOrder []int) []int {
	outCityList := make([]int, len(citySelectionOrder))
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

	//img:= image.NewRGBA(image.Rect(0,0, maxCoordinate, maxCoordinate))
	//for _, city := range cityList  {
	//	img.Set(city.X, city.Y, color.RGBA{0x88,0xff,0x88,0xff})
	//}
	//file, err := os.Create("simple.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//jpeg.Encode(file, img, &jpeg.Options{80})


	// calculate distances
	distances := calculateDistances(cityList)
	fmt.Println("Distances")
	fmt.Println(distances)

	// population
	population:=generateInitalPopulation()
	fmt.Println("Population")
	for _, item := range population	{
		fmt.Println(*item)
	}

	// estimation
	for _, item:=range population  {
		item.Estimate(cityList, distances)
	}

	// selection
	newPopulation:=make([]*Route, len(population))
	for i:=0; i< populationLength; i++ {
		item := selectFromPopulation(population)
		newPopulation[i] = item
	}

	fmt.Println("new population")
	for _, item := range newPopulation	{
		fmt.Println(*item)
	}
}

func selectFromPopulation(population []*Route) *Route {
	const tourneySize = 10
	var bestRoute *Route
	for i := 0; i < tourneySize; i++ {
		itemIndex := rand.Intn(len(population))
		selectedItem := population[itemIndex]
		if bestRoute == nil || selectedItem.Length < bestRoute.Length {
			bestRoute = selectedItem
		}
	}

	return bestRoute
}

func calculateDistances(cityList []City) [][]float64 {
	distances := make([][]float64, len(cityList))
	for i, cityFrom := range cityList {
		distances[i] = make([]float64, len(cityList))
		for _, cityTo := range cityList {
			distance := float64(0)
			if cityFrom.Id != cityTo.Id {
				squareSum := math.Pow(float64(cityFrom.X-cityTo.X), 2)+math.Pow(float64(cityFrom.Y-cityTo.Y), 2)
				distance = math.Sqrt(squareSum)
			}

			distances[cityFrom.Id][cityTo.Id] = distance
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

func generateInitalPopulation() ([]*Route) {
	population := make([]*Route, populationLength)
	for i := 0; i < populationLength; i++ {
		citySelectionOrder :=generateInitalSelectionOrder()
		population[i] = &Route{ citySelectionOrder, 0}
	}

	return population
}
