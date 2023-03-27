package tipoStruct

import "fmt"

type Estructura struct {
	numero        int
	cadena        string
	variosNumeros []int
}

func crearcionEstructura() Estructura {
	var estructura Estructura

	estructura.cadena = "Esto es una cadena"
	estructura.numero = 11
	estructura.variosNumeros = []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	return estructura
}

func RecorrerEstructura() {
	estructura := crearcionEstructura()
	fmt.Println("Cadena:: ", estructura.cadena)
	fmt.Println("Número:: ", estructura.numero)
	fmt.Println("Varios Números:: ", estructura.variosNumeros)

}
