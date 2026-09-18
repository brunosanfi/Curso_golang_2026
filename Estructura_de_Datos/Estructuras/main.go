package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Lionel"
	p.edad = 40
	p.correo = "lionel@futbol.com"
	fmt.Println(p)

	p2 := Persona{
		nombre: "Cristiano",
		edad:   39,
		correo: "cristiano@futbol.com",
	}
	fmt.Println(p2)
}
