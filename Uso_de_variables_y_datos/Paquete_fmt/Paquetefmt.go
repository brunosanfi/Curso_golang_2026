package main

import "fmt"

func main() {
	fmt.Println("Hola mundo!!")
	fmt.Print("Otro mensaje!!")
	fmt.Println("Y aca el tercero")

	name := "Lionel"
	age := 40
	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name, age)
	fmt.Println(greeting)

	var name2, lastName string
	var age2 int

	fmt.Println("Ingrese su nombre y su apellido:")
	fmt.Scanln(&name2, &lastName)
	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&age2)

	fmt.Printf("Usted se llama %s %s y tiene %d años.\n", name2, lastName, age2)

	fmt.Printf("El tipo de dato de la edad es %T\n", age2)
}
