// Enunciado: Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.

package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64

	fmt.Println("Ingrese el primer lado del triangulo:")
	fmt.Scanln(&lado1)
	fmt.Println("Ingrese el segundo lado del triangulo:")
	fmt.Scanln(&lado2)

	var perimetro, area float64

	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro = lado1 + lado2 + hipotenusa
	area = (lado1 * lado2) / 2

	fmt.Printf("El perimetro del triangulo es: %.2f\n", perimetro)
	fmt.Printf("El area del triangulo es: %.2f\n", area)
}
