package main

import "fmt"

func main() {
	//declaración de constantes

	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("p1", pi)
	fmt.Println("p2", pi2)

	//declaración de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	//area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

}
