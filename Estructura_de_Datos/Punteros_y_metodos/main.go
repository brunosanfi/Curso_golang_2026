package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var x int = 10

	var p *int = &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Dirección de memoria de x:", &x)
	fmt.Println("Valor de p (dirección de memoria de x):", p)
	fmt.Println("Valor al que apunta p (valor de x):", *p)

	editar(&x)
	fmt.Println("Valor de x:", x)

	p1 := Persona{
		nombre: "Lionel",
		edad:   40,
		correo: "lionel@futbol.com",
	}
	p1.saludar()

}

func editar(valor *int) {
	*valor = 20
	fmt.Println("Valor editado:", *valor)
}

func (p *Persona) saludar() {
	fmt.Printf("Hola, me llamo %s y tengo %d años. Mi correo es %s\n", p.nombre, p.edad, p.correo)
}
