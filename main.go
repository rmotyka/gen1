package main

import (
	"fmt"
	"math/rand"
	//"image"
	//"image/color"
	//"os"
	//"log"
	//"image/jpeg"
	//"time"
	"gen1/structs"
	"gen1/selection"
	"gen1/initialize"
)


const populationLength = 100
const numberOfCities = 10
const maxCoordinate = 100

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("gen2")

	cityList := initialize.GetCityList(numberOfCities, maxCoordinate)
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
	distances := initialize.CalculateDistances(cityList)
	fmt.Println("Distances")
	fmt.Println(distances)

	// population
	population:=generateInitalPopulation()
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
		crossing(newPopulation)

		for _, item := range newPopulation {
			item.Estimate(cityList, distances)
		}

		fmt.Println("after crossing")
		for _, item := range newPopulation {
			fmt.Println(*item)
		}

		// mutation
		mutate(newPopulation)

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

func mutate(population []*structs.Route) {
	numberOfMutations := int(populationLength * 1/100);
	for i :=0; i<numberOfMutations; i++ {
		itemIndex := rand.Intn(len(population))
		item := population[itemIndex]

		cityIndex := rand.Intn(numberOfCities)
		newCityIndex := rand.Intn(numberOfCities - cityIndex)
		item.CitySelectionOrder[cityIndex] = newCityIndex
		item.Length = 0
	}
}

func crossing(population []*structs.Route) {
	numberOfCrossing := int(populationLength * 10/100);
	for i :=0; i<numberOfCrossing; i++ {
		parentAIndex := rand.Intn(len(population))
		parentBIndex := rand.Intn(len(population))
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

func generateInitalSelectionOrder() []int {
	citySelectionOrder:=make([]int, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		citySelectionOrder[i] = rand.Intn(numberOfCities-i)
	}

	return citySelectionOrder
}

func generateInitalPopulation() ([]*structs.Route) {
	population := make([]*structs.Route, populationLength)
	for i := 0; i < populationLength; i++ {
		citySelectionOrder :=generateInitalSelectionOrder()
		population[i] = &structs.Route{ citySelectionOrder, 0}
	}

	return population
}
