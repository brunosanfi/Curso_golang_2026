package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#ff0000",
		"verde": "#00ff00",
		"azul":  "#0000ff",
	}

	fmt.Println(colors)

	fmt.Println("El color rojo es:", colors["rojo"])
	fmt.Println("El color verde es:", colors["verde"])
	fmt.Println("El color azul es:", colors["azul"])

	colors["amarillo"] = "#ffff00"
	fmt.Println("El color amarillo es:", colors["amarillo"])

	fmt.Println("Vamos a probar el verificador:")

	if color, ok := colors["negro"]; ok {
		fmt.Println("El color negro es:", color)
	} else {
		fmt.Println("El color negro no existe en el mapa")
	}

	delete(colors, "verde")
	fmt.Println("El color verde ha sido eliminado del mapa")

	for clave, valor := range colors {
		fmt.Println("El color", clave, "tiene el valor", valor)
	}
}
