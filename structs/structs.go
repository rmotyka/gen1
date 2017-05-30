package structs

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
	//fmt.Println("City in Route")
	//fmt.Println(citiesInRouteOrder)

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
	//fmt.Println("Route lenght ", routeLenght)
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

type ByLenght []*Route

func (s ByLenght) Len() int {
	return len(s)
}
func (s ByLenght) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLenght) Less(i, j int) bool {
	return s[i].Length < s[j].Length
}
