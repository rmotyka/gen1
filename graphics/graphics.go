package graphics

import (
	"github.com/ajstarks/svgo"
	"os"
	"gen1/structs"
	"strconv"
)

func SaveCityImage(name string, maxCoordinate int, cityList []structs.City) {
	width := maxCoordinate * 10
	height := maxCoordinate * 10

	f, _ := os.Create( name + ".svg")
	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(width, height)

	var previousCity structs.City
	for _, city := range cityList {
		canvas.Circle(city.X * 10, city.Y * 10, 2)
		canvas.Text(city.X * 10, city.Y * 10, strconv.Itoa(city.Id), "text-anchor:middle;font-size:50px;fill:red")
		if previousCity.X != 0 && previousCity.Y != 0 {
			canvas.Line(previousCity.X*10, previousCity.Y*10, city.X*10, city.Y*10, "stroke:#006600;")
		}

		previousCity = city
	}

	canvas.End()
}
