package geometry
import (
	"math"
)
// Point is a point
type Circle struct {
	radio float64
	angle float64
}

// MakePoint is a point function
func MakeCircle(r, a float64) Circle {
	p1 := Circle{r, a} 
	return p1
}

func (obj* Circle) Getx() float64{
	return  obj.radio*math.Sin(obj.angle)
}

func (obj* Circle) Gety() float64{
	return  obj.radio*math.Cos(obj.angle)
}

func (obj* Circle) Setangle(newangle float64) {
	obj.angle = newangle
}