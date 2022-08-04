package geometric

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) AreaAndPerimeter() (float64, float64) {
	return math.Pi * (c.Radius * c.Radius), 2 * math.Pi * c.Radius
}
