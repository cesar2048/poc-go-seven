package main

import (
	"fmt"
	"math"
	"ifcommit.com/seven/geometry"
)

func main() {
	fmt.Println("Hola mundo")
	p1 := geometry.MakePoint(10, 25)
	/*using function Point*/
	/*p2 := p1.x * 2
	fmt.Println(p1.x)
	fmt.Println(p2)*/
	fmt.Println(p1.Getx())
	/*using function circulo (MakeCircle) and methods ..*/
	circulo := geometry.MakeCircle(5,(math.Pi/4))
	fmt.Println(circulo.Getx())
	fmt.Println(circulo.Gety())
	circulo.Setangle(math.Pi/2) //change a 90 degre-->rad 
	fmt.Println(circulo.Getx())
	fmt.Println(circulo.Gety())
}
