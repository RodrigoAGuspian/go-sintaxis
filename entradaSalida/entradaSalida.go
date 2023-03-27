package entradaSalida

import (
	"fmt"
	"strconv"
)

func TiposDeSalida() {
	imprimo := "imprimo"
	numero := 1
	fmt.Print("Solo\n", imprimo, numero)
	fmt.Printf("Solo %s como C %d\n", imprimo, numero)
	fmt.Println("Solo ", imprimo, numero, " Pero con renglón :)")

}

func TiposDeEntrada() {
	receptor := ""
	fmt.Print("Ingrese una palabra (fmt.Scan): ")
	fmt.Scan(receptor)
	fmt.Println(receptor)

	fmt.Print("Ingrese otra palabra (fmt.Scanf)")
	fmt.Scanf(receptor)
	fmt.Println(receptor)

	fmt.Print("Ingrese otra palabra (fmt.Scanln)")
	fmt.Scanln(receptor)
	fmt.Println(receptor)

}

func IngresarNumeros() {
	// Mi recomendación.
	receptor := ""

	fmt.Print("Ingrese un número")
	fmt.Scanln(receptor)

	numero, err := strconv.Atoi(receptor)

	if err != nil {
		panic("Erorr de conversión de entero.")
	} else {
		fmt.Println("El número ingresado es: ", numero)
	}

}

func IngresarNumeroPositivo() int {
	// Mi recomendación.
	receptor := ""
	for {

		fmt.Print("Ingrese un número")
		fmt.Scanln(receptor)

		numero, err := strconv.Atoi(receptor)
		if err != nil {
			panic("Error de conversión de entero.")
		} else {
			return numero
		}

	}

}
