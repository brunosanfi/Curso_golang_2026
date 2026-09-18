package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	titulo      string
	descripcion string
	completado  bool
}

type ListasDeTareas struct {
	tareas []Tarea
}

func (l *ListasDeTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *ListasDeTareas) marcarCompletado(indice int) {
	if indice >= 0 && indice < len(l.tareas) {
		l.tareas[indice].completado = true
	}
}

func (l *ListasDeTareas) editarTarea(indice int, t Tarea) {
	if indice >= 0 && indice < len(l.tareas) {
		l.tareas[indice] = t
	}
}

func (l *ListasDeTareas) eliminarTarea(indice int) {
	if indice >= 0 && indice < len(l.tareas) {
		l.tareas = append(l.tareas[:indice], l.tareas[indice+1:]...)
	}
}

func main() {
	lista := ListasDeTareas{}

	leer := bufio.NewReader((os.Stdin))

	for {
		var opcion int

		fmt.Println("Seleccione una opción\n",
			"1.Agregar Tarea\n",
			"2.Marcar Tarea como completada\n",
			"3.Editar Tarea\n",
			"4.Eliminar Tarea\n",
			"5.Listar Tareas\n",
			"6.Salir",
		)

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la Tarea: ")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la Tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea completada:")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el índice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Print("Ingrese el descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
		case 5:
			fmt.Println("Lista de tareas:")
			fmt.Println("=======================================================================================")
			for i, t := range lista.tareas {
				fmt.Printf("%d. %s - %s (Completado: %t)\n", i+1, t.titulo, t.descripcion, t.completado)
			}
			fmt.Println("=======================================================================================")

		case 6:
			fmt.Print("Saliendo del programa...")
			return
		default:
			fmt.Print("Opcion invalida")
		}

	}
}
