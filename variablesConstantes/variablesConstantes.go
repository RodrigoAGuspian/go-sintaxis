package variablesConstantes

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
  Declaración de constantes
  const constante = valor
*/

const NUM_PROG = 2
const NOM_PROG = "main.go"

const (
	NUM_PROG1 = 1
	NOM_PROG1 = "main.go"
)

/*
  Declaración de variables
  var variable tipo = valor
  variable := valor
*/

func Booleanos() {
	var bandera bool = true
	bandera1 := false
	fmt.Println(bandera, bandera1)

}

func Enteros() {
	var entero int = 0
	entero1 := 1
	fmt.Println(entero, entero1)
}

func Flotantes() {
	var decimal float32 = 3.14
	decimal1 := 3.15
	fmt.Println("Decimales: ", decimal, decimal1,
		reflect.TypeOf(decimal), reflect.TypeOf(decimal1))
}

func Complejos() {
	var ai complex64 = complex(1, 6)
	bi := complex(2, 3)

	fmt.Println(ai, bi)

	fmt.Println("Parte real:", real(ai))
	fmt.Println("Parte imaginaria:", imag(ai))
}

func Caracteres() {
	var caracter byte = 'a'
	var caracterB rune = 'b'
	var caracterX rune = 120
	caracterC := 'c'

	fmt.Printf("%c, %c, %c, %c \n", caracter, caracterB, caracterC, caracterX)

}

func Cadenas() {
	var cadena string = "Cadena de caracteres"
	cadena1 := "Cadena de caracteres 1"
	fmt.Printf("%s\n%s\n", cadena, cadena1)

	fmt.Println("Concatenación", cadena+cadena1)
}

func ConvertirCadenaAEntero() {
	s := "123"

	// Cadena a entero
	i, err := strconv.Atoi(s)
	if err != nil {
		// Error por si sale mal la conversión
		panic(err)
	}

	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))

}

func ConvertirCadenaAFloat() {
	cadena := "3.14159265"
	if s, err := strconv.ParseFloat(cadena, 32); err == nil {
		fmt.Println(s, reflect.TypeOf(s))
		fmt.Println(cadena, reflect.TypeOf(cadena))
	}
	if s, err := strconv.ParseFloat(cadena, 64); err == nil {
		fmt.Println(s, reflect.TypeOf(s))
		fmt.Println(cadena, reflect.TypeOf(cadena))
	}
}

func ConvertirEnteroACadena() {
	i := 1000
	// Entero a cadena
	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)

	fmt.Println(s1, reflect.TypeOf(s1))
	fmt.Println(s2, reflect.TypeOf(s2))
	fmt.Println(i, reflect.TypeOf(i))
}

func ConvetirFloatACadena() {
	f := 123.456
	s := fmt.Sprintf("%f", f)

	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(s, reflect.TypeOf(s))

}

func ConvertirFloatAEntero() {
	f := 3.12
	a := int64(f)

	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(a, reflect.TypeOf(a))
}

func ConvertirEnteroAFloat() {
	a := 312
	f := float64(a)

	fmt.Println(f, reflect.TypeOf(f))
	fmt.Println(a, reflect.TypeOf(a))
}
