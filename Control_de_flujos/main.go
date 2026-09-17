package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/*if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Es de mañana!!!")
	} else if t.Hour() < 18 {
		fmt.Println("Es de tarde!!!")
	} else {
		fmt.Println("Es de noche!!!")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Estás en Windows")
	case "linux":
		fmt.Println("Estás en Linux")
	case "darwin":
		fmt.Println("Estás en macOS")
	default:
		fmt.Println("Sistema operativo no reconocido")
	}

	for i := 1; i <= 12; i++ {
		fmt.Println(i)

		if i == 5 {
			continue // Saltar a la siguiente iteración cuando i sea igual a 5
		}

		if i == 10 {
			break // Salir del bucle cuando i sea igual a 10
		}
	}

	saludo := hello("Bruno") // Llamada a la función anónima
	fmt.Println(saludo)

	suma, multi := calc(5, 3) // Llamada a la función con múltiples valores de retorno
	fmt.Println("La suma es:", suma)
	fmt.Println("La multiplicación es:", multi)*/

	jugar()

}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número!")
			jugarNuvamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor.")
		}
	}
	fmt.Println("Lo siento, no adivinaste el número. El número era:", numAleatorio)
	jugarNuvamente()
}

func jugarNuvamente() {
	var eleccion string
	fmt.Print("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Elección inválida. Inténtalo nuevamente.")
		jugarNuvamente()
	}
}

/*
func hello(name string) string {
	return "Hola " + name + " desde una función anónima!"
}

func calc(a, b int) (sum, mult int) {
	sum = a + b
	mult = a * b

	return sum, mult
}
*/
