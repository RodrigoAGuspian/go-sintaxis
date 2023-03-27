package funcionesProcedimientos

import "fmt"

/*
	Declaración de funciones
	func nombreFuncion( [parámetro Tipo], [...] ) tipoARetornar{
		return valorRetornar
	}

*/

func FuncionMultiplicacion(x int) int {
	valor := 10 * x
	return valor
}

/*
	Declaración de procedimientos
	func nombreProcedimiento( [parametro Tipo], [...] ){
	}

*/
func ProcedimientoDivision(x int) {
	if x == 0 {
		fmt.Println("No se puede Dividir entre 0")
	}

	valor := 10 / x

	fmt.Println("La división entre 10 y", x, "es: ", valor)

}

/* Se denota el primer carácter en minúscula para decir que es privado
y si es mayúscula se dice que es público
*/

func funcionPrivada() string {
	return "Soy una función privada"
}

func FuncionPublica() string {
	return "Soy una función público"
}

func procedimientoPrivado() {
	fmt.Println("Soy un procedimiento privado")
}

func ProcedimientoPublico() {
	fmt.Println("Soy un procedimiento público")
}

/* Sí el parámetro es el mismo se puede construir de la siguiente manera:
func nombreFuncion( [parámetro1, parámetro2 Tipo])tipoARetornar{
	return valorRetornar
}
*/

/* Para devolver más valores se hace de la siguiente manera
func nombreFuncion( [parámetro1, parámetro2 Tipo])(tipoARetornar1, tipoARetornar2){
	return valorRetornar1, valorRetornar2
}
*/
func SumaYResta(i, j int) (int, int) {
	return i + j, i * j
}
