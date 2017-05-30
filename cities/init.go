package cities

import (
	"math/rand"
	"gen1/structs"
	"math"
)

func GetCityList(numberOfCities int, maxCoordinate int) ([]structs.City) {
	cityList := make([]structs.City, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		cityList[i] = structs.City{ i,rand.Intn(maxCoordinate),rand.Intn(maxCoordinate)}
	}

	return cityList;
}


func CalculateDistances(cityList []structs.City) [][]float64 {
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