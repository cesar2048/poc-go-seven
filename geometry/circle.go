package geometry
import (
	"math"
	"fmt"
)
// Point is a point
type Circle struct {
	radio float64
	angle float64
}

// MakePoint is a point function
func MakeCircle(r, a float64) Circle {
	newcircle := Circle{r, a} 
	return newcircle
}
/*funcion en donde recibimos obj por referencia (usando punteros)
func (obj* Circle) Setangle(newangle* float64) {
	fmt.Println("entro setangle")
	//obj.angle = newangle
	fmt.Printf("1. en Setangle   &obj.angle=%p \n",&obj.angle)
	*newangle= 5
}*/

func (obj* Circle) Setangle(newangle float64) {
	fmt.Println("entro setangle")
	obj.angle = newangle
	fmt.Printf("1. en Setangle   &obj.angle=%p \n",&obj.angle)
}

func (obj* Circle) Getx() float64{
	fmt.Printf("1. en getx   &obj.angle=%p \n",&obj.angle)
	return  obj.radio*math.Sin(obj.angle)
}

func (obj* Circle) Gety() float64{
	return  obj.radio*math.Cos(obj.angle)
}

