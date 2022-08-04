package geometric

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) AreaAndPerimeter() (float64, float64) {
	return r.Width * r.Height, r.Width*2 + r.Height*2
}
