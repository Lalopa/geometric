package geometric

import "fmt"

type AreaAndPerimeter interface {
	AreaAndPerimeter() (float64, float64)
}

func GetAreaAndPerimeter(areaAndPerimeter AreaAndPerimeter) {
	area, perimeter := areaAndPerimeter.AreaAndPerimeter()
	fmt.Println(areaAndPerimeter)
	fmt.Printf("Area: %f, Perimetro: %f\n", area, perimeter)
}
