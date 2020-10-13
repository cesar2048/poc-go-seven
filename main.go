package main

import (
	"fmt"

	"ifcommit.com/seven/geometry"
)

func main() {
	fmt.Println("Hola mundo")
	p1 := geometry.MakePoint(1, 2)
	fmt.Print(p1.x)
}
