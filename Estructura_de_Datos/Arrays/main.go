package main

import "fmt"

func main() {
	var matriz [5]int

	fmt.Println(matriz)

	matriz[0] = 10
	fmt.Println(matriz)

	var matriz2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(matriz2)

	var matriz3 = [...]int{15, 25, 35, 45, 55}
	fmt.Println(matriz3)

	println("Recorro con el bucle for")

	for i := 0; i < len(matriz3); i++ {
		fmt.Println("El valor de la posicion", i, "es", matriz3[i])
	}

	println("Ahora con el bucle for range")

	for index, value := range matriz3 {
		fmt.Println("El valor de la posicion", index, "es", value)
	}

	var matrizMulti = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Matriz multidimensional:", matrizMulti)

}
