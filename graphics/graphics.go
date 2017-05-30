package graphics

import (
	"github.com/ajstarks/svgo"
	"os"
	"gen1/structs"
)

func SaveCityImage(maxCoordinate int, cityList []structs.City) {
	width := 500
	height := 500

	f, _ := os.Create("outsvg.svg")
	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()


}
