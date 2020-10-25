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
	circulo := geometry.MakeCircle(5,math.Pi/2)
	fmt.Println(circulo.Getx())
	fmt.Println(circulo.Gety())
	/*INICIO --- valores en main usando  funcion en donde recibimos obj por referencia (usando punteros)
	angulotmp := (math.Pi/4)
	circulo.Setangle(&angulotmp) //change a 90 degre-->rad 
	fmt.Printf("angulotmp %v",angulotmp) ---FIN*/
	angulotmp := (math.Pi/4)
	circulo.Setangle(angulotmp) //change a 90 degre-->rad 
	fmt.Println("nuevos valores con angulo cambiado")
	fmt.Println(circulo.Getx())
	fmt.Println(circulo.Gety())
}
