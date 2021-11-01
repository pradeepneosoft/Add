package Add

// if center is 0,0
type Init struct {
	A, B float64
}

// Get Eccentricity of Ellipse
func (e *Init) Add() float64 {
	return e.A + e.B
}
