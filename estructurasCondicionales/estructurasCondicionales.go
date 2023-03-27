package estructurasCondicionales

import (
	"fmt"
	"main/entradaSalida"
)

func EstructuraDecision() {
	num := entradaSalida.IngresarNumeroPositivo()
	if num < 0 {
		fmt.Println(num, "es negativo")
	} else if num < 10 {
		fmt.Println(num, "tiene un digito")
	} else {
		fmt.Println(num, "tiene muchos digitos")
	}

}

func EstructuraSeleccion() {
	genero := 'f'
	switch genero {
	case 'm':
		fmt.Println("Masculino")
	case 'f':
		fmt.Println("Femenino")
	}
}

func EsctructuraSeleccionTradicional() {
	desde := 1
	cadenaNumeros := ""
	switch desde {
	case 1:
		cadenaNumeros += "1, "
		fallthrough
	case 2:
		cadenaNumeros += "2, "
		fallthrough
	case 3:
		cadenaNumeros += "3, "
		fallthrough
	case 4:
		cadenaNumeros += "4, "
		fallthrough
	case 5:
		cadenaNumeros += "5, "
	}

	fmt.Println(cadenaNumeros)
}
