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

func (r *Route) Equals(other *Route) bool {
	a := r.CitySelectionOrder
	b := other.CitySelectionOrder

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}

func (r *Route) Estimate(cityList []City, distances [][]float64) {
	citiesInRouteOrder :=SelectCitiesIds(cityList, r.CitySelectionOrder)
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


func SelectCitiesIds(cityList []City, citySelectionOrder []int) []int {
	tempCityList := make([]City, len(cityList))
	copy(tempCityList, cityList)

	outCityList := make([]int, len(citySelectionOrder))
	for i:=0; i<len(citySelectionOrder); i++ {
		indexOfCity := citySelectionOrder[i]
		city := tempCityList[indexOfCity]
		outCityList[i] = city.Id

		// remove indexOfCity from cityList:
		tempCityList = append(tempCityList[:indexOfCity], tempCityList[indexOfCity+1:]...)
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
