package geometry

// Point is a point
type Point struct {
	x float64
	y float64
}

// MakePoint is a point function
func MakePoint(x, y float64) Point {
	p1 := Point{x, y} 
	return p1
}

func (mi Point) Getx() float64{
	return mi.x
}
