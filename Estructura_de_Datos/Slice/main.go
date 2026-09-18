package main

import "fmt"

func main() {
	var dias = []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	fmt.Println("Los dias de la semana son:")

	for i := 0; i < len(dias); i++ {
		fmt.Println(dias[i])
	}

	diasDeSemana := dias[1:6]
	fmt.Println("Los dias de la semana son:", diasDeSemana)

	fmt.Println(len(diasDeSemana))
	fmt.Println(cap(diasDeSemana))

	println("Agregamos un nuevo dia a la semana")
	diasDeSemana = append(diasDeSemana, "Feriado")

	println("Eliminamos un dia de la semana:")
	diasDeSemana = append(diasDeSemana[:2], diasDeSemana[3:]...)

	sliceEntero := []int{1, 2, 3, 4, 5}
	sliceEnteroMake := make([]int, 5)

	fmt.Println("Asi están inicialmente los arrays de enteros:")

	fmt.Println("SliceEntero:", sliceEntero)
	fmt.Println("SliceEnteroMake:", sliceEnteroMake)

	fmt.Println("Copiamos el sliceEntero al sliceEnteroMake")

	copy(sliceEnteroMake, sliceEntero)

	fmt.Println("Después de la copia:")
	fmt.Println("SliceEntero:", sliceEntero)
	fmt.Println("SliceEnteroMake:", sliceEnteroMake)

}
