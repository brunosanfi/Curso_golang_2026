package main

import (
	"fmt"
	"math"
)

func main() {
	// Operadores aritméticos básicos
	a := 10
	b := 3

	suma := a + b
	resta := a - b
	multiplicacion := a * b
	division := a / b
	modulo := a % b

	fmt.Println("Operadores aritméticos:")
	fmt.Println("a =", a, "b =", b)
	fmt.Println("suma:", suma)
	fmt.Println("resta:", resta)
	fmt.Println("multiplicacion:", multiplicacion)
	fmt.Println("division:", division)
	fmt.Println("modulo:", modulo)

	// Operadores de incremento y decremento
	fmt.Println("\nIncremento y decremento:")
	contador := 5
	contador++
	fmt.Println("contador++ =>", contador)
	contador--
	fmt.Println("contador-- =>", contador)

	// Operador de asignación
	fmt.Println("\nOperador de asignación:")
	valor := 20
	valor = 30
	fmt.Println("valor = 30 =>", valor)

	// Asignación combinada con operadores aritméticos
	fmt.Println("\nAsignación combinada:")
	x := 8
	x += 2
	fmt.Println("x += 2 =>", x)
	x -= 3
	fmt.Println("x -= 3 =>", x)
	x *= 4
	fmt.Println("x *= 4 =>", x)
	x /= 2
	fmt.Println("x /= 2 =>", x)
	x %= 3
	fmt.Println("x %= 3 =>", x)

	// Funciones principales del paquete math
	fmt.Println("\nFunciones del paquete math:")
	fmt.Println("math.Abs(-7.5) =", math.Abs(-7.5))           // devuelve el valor absoluto
	fmt.Println("math.Pow(2, 3) =", math.Pow(2, 3))           // calcula una potencia
	fmt.Println("math.Sqrt(81) =", math.Sqrt(81))             // calcula la raíz cuadrada
	fmt.Println("math.Ceil(4.2) =", math.Ceil(4.2))           // redondea hacia arriba
	fmt.Println("math.Floor(4.9) =", math.Floor(4.9))         // redondea hacia abajo
	fmt.Println("math.Round(4.6) =", math.Round(4.6))         // redondea al entero más cercano
	fmt.Println("math.Max(10, 20) =", math.Max(10, 20))       // devuelve el mayor de dos valores
	fmt.Println("math.Min(10, 20) =", math.Min(10, 20))       // devuelve el menor de dos valores
	fmt.Println("math.Pi =", math.Pi)                         // constante del valor pi
	fmt.Println("math.E =", math.E)                           // constante del número e
	fmt.Println("math.Sin(math.Pi/2) =", math.Sin(math.Pi/2)) // calcula el seno
	fmt.Println("math.Cos(0) =", math.Cos(0))                 // calcula el coseno
}
