package main

import "fmt"

type City struct {
	X int
	Y int
}

type Route struct {
	CitySelectionOrder []int
}

func getCityList() (map[int]City) {
	cityList := make(map[int]City)
	cityList[0] = City{0,0}
	cityList[1] = City{1,1}
	cityList[3] = City{3,5}

	return cityList;
}

func main() {
	fmt.Println("gen2")

	//cityList := getCityList()
	population:=generateInitalPopulation()

}
func generateInitalPopulation() ([]Route) {

}
