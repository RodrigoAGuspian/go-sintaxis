package entradaSalida

import (
	"fmt"
	"strconv"
)

func TiposDeSalida() {
	imprimo := "imprimo"
	numero := 1
	fmt.Print("Solo ", imprimo, " ", numero, " \n")
	fmt.Printf("Solo %s como C %d\n", imprimo, numero)
	fmt.Println("Solo", imprimo, numero, "Pero con renglón :)")

}

func TiposDeEntrada() {
	receptor := ""

	fmt.Println("Ingrese una palabra (fmt.Scanf): ")
	fmt.Scanf("%s", &receptor)
	fmt.Println(receptor)

	fmt.Println("Ingrese otra palabra (fmt.Scan): ")
	fmt.Scan(&receptor)
	fmt.Println(receptor)

	fmt.Scanln() // Tratar de hacer una limpieza en la entrada estandar

	fmt.Println("Ingrese otra palabra (fmt.Scanln): ")
	fmt.Scanln(&receptor)
	fmt.Println(receptor)

}

func IngresarNumeros() {
	// Mi recomendación.
	receptor := ""

	fmt.Print("Ingrese un número")
	fmt.Scanln(&receptor)

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

		fmt.Print("Ingrese un número: ")
		fmt.Scanln(&receptor)

		numero, err := strconv.Atoi(receptor)
		if err != nil {
			panic("Error de conversión de entero.")
		} else {
			return numero
		}

	}

}
