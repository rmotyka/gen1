package graphics

import (
	"github.com/ajstarks/svgo"
	"os"
	"gen1/structs"
)

func SaveCityImage(maxCoordinate int, cityList []structs.City) {
	width := maxCoordinate
	height := maxCoordinate

	f, _ := os.Create("outsvg.svg")
	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(width, height)

	for _, city := range cityList {
		canvas.Circle(city.X, city.Y, 2)
	}


	//canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()


}
