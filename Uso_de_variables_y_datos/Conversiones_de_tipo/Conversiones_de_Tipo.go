package main

import (
	"fmt"
	"strconv"
)

func main() {
	var variableInt16 int16 = 50
	var variableInt32 int32 = 100

	s := "100"

	i, _ := strconv.Atoi(s)

	fmt.Println("La suma de los dos numeros es:", int32(variableInt16)+variableInt32)
	fmt.Println("El doble del valor convertido es:", i*2)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println("El valor convertido a string y concaternado es:", s+s)
}
