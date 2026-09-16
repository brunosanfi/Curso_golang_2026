package main

import (
	"fmt"
	"math"
)

func main() {
	// Enteros con signo
	var enteroInt int = 42
	var enteroInt8 int8 = 127
	var enteroInt16 int16 = 32767
	var enteroInt32 int32 = 2147483647
	var enteroInt64 int64 = 9223372036854775807

	// Enteros sin signo
	var sinSignoUint uint = 42
	var sinSignoUint8 uint8 = 255
	var sinSignoUint16 uint16 = 65535
	var sinSignoUint32 uint32 = 4294967295
	var sinSignoUint64 uint64 = 18446744073709551615

	// Flotantes
	var decimal32 float32 = 3.14159265
	var decimal64 float64 = 3.141592653589793

	fmt.Println("Enteros con signo:")
	fmt.Printf("int     => %T = %d\n", enteroInt, enteroInt)
	fmt.Printf("int8    => %T = %d\n", enteroInt8, enteroInt8)
	fmt.Printf("int16   => %T = %d\n", enteroInt16, enteroInt16)
	fmt.Printf("int32   => %T = %d\n", enteroInt32, enteroInt32)
	fmt.Printf("int64   => %T = %d\n", enteroInt64, enteroInt64)
	fmt.Printf("min int => %d\n", int(math.MinInt))
	fmt.Printf("max int => %d\n", int(math.MaxInt))
	fmt.Printf("min int8 => %d\n", math.MinInt8)
	fmt.Printf("max int8 => %d\n", math.MaxInt8)
	fmt.Printf("min int16 => %d\n", math.MinInt16)
	fmt.Printf("max int16 => %d\n", math.MaxInt16)
	fmt.Printf("min int32 => %d\n", math.MinInt32)
	fmt.Printf("max int32 => %d\n", math.MaxInt32)
	fmt.Printf("min int64 => %d\n", math.MinInt64)
	fmt.Printf("max int64 => %d\n", math.MaxInt64)

	fmt.Println("\nEnteros sin signo:")
	fmt.Printf("uint    => %T = %d\n", sinSignoUint, sinSignoUint)
	fmt.Printf("uint8   => %T = %d\n", sinSignoUint8, sinSignoUint8)
	fmt.Printf("uint16  => %T = %d\n", sinSignoUint16, sinSignoUint16)
	fmt.Printf("uint32  => %T = %d\n", sinSignoUint32, sinSignoUint32)
	fmt.Printf("uint64  => %T = %d\n", sinSignoUint64, sinSignoUint64)
	fmt.Printf("min uint => %d\n", 0)
	fmt.Printf("max uint => %d\n", uint(math.MaxUint))
	fmt.Printf("min uint8 => %d\n", 0)
	fmt.Printf("max uint8 => %d\n", math.MaxUint8)
	fmt.Printf("min uint16 => %d\n", 0)
	fmt.Printf("max uint16 => %d\n", math.MaxUint16)
	fmt.Printf("min uint32 => %d\n", 0)
	fmt.Printf("max uint32 => %d\n", math.MaxUint32)
	fmt.Printf("min uint64 => %d\n", 0)
	fmt.Printf("max uint64 => %d\n", uint64(math.MaxUint64))

	fmt.Println("\nTipos flotantes:")
	fmt.Printf("float32 => %T = %.8f\n", decimal32, decimal32)
	fmt.Printf("float64 => %T = %.15f\n", decimal64, decimal64)
	fmt.Printf("min float32 => %e\n", math.SmallestNonzeroFloat32)
	fmt.Printf("max float32 => %e\n", math.MaxFloat32)
	fmt.Printf("min float64 => %e\n", math.SmallestNonzeroFloat64)
	fmt.Printf("max float64 => %e\n", math.MaxFloat64)

	fmt.Println("\nNota:")
	fmt.Println("- int y uint cambian según la arquitectura del sistema.")
	fmt.Println("- float32 tiene menos precisión que float64.")
	fmt.Println("- En la mayoría de casos se usa float64 para números decimales.")

	// Booleanos
	var variableBolleana bool = true
	fmt.Println("\nBooleanos:")
	fmt.Printf("bool => %T = %t\n", variableBolleana, variableBolleana)

	// Strings y caracteres de escape
	fmt.Println("\nStrings y caracteres de escape:")
	fmt.Printf("\\n => salto de línea: %q\n", "\n")
	fmt.Printf("\\t => tabulación: %q\n", "\t")
	fmt.Printf("\\r => retorno de carro: %q\n", "\r")
	fmt.Printf("\\b => retroceso: %q\n", "\b")
	fmt.Printf("\\f => avance de página: %q\n", "\f")
	fmt.Printf("\\' => comilla simple: %q\n", "'")
	fmt.Printf("\\\" => comilla doble: %q\n", "\"")
	fmt.Printf("\\\\ => barra invertida: %q\n", "\\")

	var saludo = "Hola\tMundo\n\tBienvenido a Go"
	fmt.Println("saludo =>", saludo)

	var ruta = "C:\\Users\\Usuario\\Documents\\archivo.txt"
	fmt.Println("ruta =>", ruta)

	var mensaje = "Mi nombre es \"Ana\" y tengo 25 años."
	fmt.Println("mensaje =>", mensaje)

	// Byte
	fmt.Println("\nTipo byte:")
	var b byte = 'A'
	var b2 byte = 65
	var bytes = []byte{72, 73, 32, 71, 79}

	fmt.Printf("byte b = %c (valor entero: %d)\n", b, b)
	fmt.Printf("byte b2 = %c (valor entero: %d)\n", b2, b2)
	fmt.Printf("slice de bytes = %q\n", bytes)
	fmt.Printf("string desde bytes = %s\n", string(bytes))
	fmt.Printf("byte tiene alias de uint8: %T\n", b)

	// Rune
	fmt.Println("\nTipo rune:")
	var r1 rune = 'A'
	var r2 rune = 'ñ'
	var r3 rune = '中'
	var r4 rune = '😊'

	fmt.Printf("rune r1 = %c (valor entero: %d)\n", r1, r1)
	fmt.Printf("rune r2 = %c (valor entero: %d)\n", r2, r2)
	fmt.Printf("rune r3 = %c (valor entero: %d)\n", r3, r3)
	fmt.Printf("rune r4 = %c (valor entero: %d)\n", r4, r4)
	fmt.Printf("rune es alias de int32: %T\n", r1)
	fmt.Println("Los runes representan caracteres Unicode, no solo ASCII.")
}
