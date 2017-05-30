package graphics

import (
	"image"
	"image/color"
	"os"
	"log"
	"image/jpeg"
	"gen1/structs"
)

func SaveCityImage(maxCoordinate int, cityList []structs.City) {
	img:= image.NewRGBA(image.Rect(0,0, maxCoordinate, maxCoordinate))
	for _, city := range cityList  {
		img.Set(city.X, city.Y, color.RGBA{0x88,0xff,0x88,0xff})
	}
	file, err := os.Create("simple.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, img, &jpeg.Options{80})
}
