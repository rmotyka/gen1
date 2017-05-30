package main

import (
	"fmt"
	//"image"
	//"image/color"
	//"os"
	//"log"
	//"image/jpeg"
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

