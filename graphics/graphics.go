package graphics

import (
	"github.com/ajstarks/svgo"
	"os"
	"gen1/structs"
	"strconv"
)

func SaveCityImage(maxCoordinate int, cityList []structs.City) {
	width := maxCoordinate * 10
	height := maxCoordinate * 10

	f, _ := os.Create("outsvg.svg")
	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(width, height)

	var previousCity structs.City
	for i, city := range cityList {
		canvas.Circle(city.X * 10, city.Y * 10, 2)
		canvas.Text(city.X * 10, city.Y * 10, strconv.Itoa(i), "text-anchor:middle;font-size:50px;fill:red")
		canvas.Line(previousCity.X * 10, previousCity.Y * 10, city.X * 10, city.Y * 10, "stroke:#006600;")

		previousCity = city
	}

	canvas.End()
}
