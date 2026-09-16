package main

import "fmt"

const PI = 3.14

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fisrtName, lastName := "Lionel", "Messi"
	age := 39

	fmt.Println("Mi nombre es", fisrtName, lastName, "y tengo", age, "años.")
	fmt.Println("El valor de PI es", PI)

	fmt.Println("El domingo es el dia", Domingo, "de la semana")
	fmt.Println("El lunes es el dia", Lunes, "de la semana")
	fmt.Println("El martes es el dia", Martes, "de la semana")
	fmt.Println("El miercoles es el dia", Miercoles, "de la semana")
	fmt.Println("El jueves es el dia", Jueves, "de la semana")
	fmt.Println("El viernes es el dia", Viernes, "de la semana")
	fmt.Println("El sabado es el dia", Sabado, "de la semana")
}
