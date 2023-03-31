package vectoresMatrices

import "fmt"

func Vector() {
	// Vectores y matrices
	var numeros [3]int
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	fmt.Println(numeros)

	numeros2 := [3]int{1, 2, 3}
	fmt.Println(numeros2)

}

func Matriz() {
	var matrizNumeros [2][2]int
	matrizNumeros[0][0] = 1
	matrizNumeros[0][1] = 2
	matrizNumeros[1] = [2]int{3, 4}
	fmt.Println(matrizNumeros)
}

func Slice() {
	colores := [4]string{"Azul", "Rojo", "Amarillo", "Verde"}
	fmt.Println(colores)

	// Slices
	coloresSlice := colores[0:2]
	fmt.Println(coloresSlice)
	// Agregar dato
	coloresSlice = append(coloresSlice, "Blanco")
	fmt.Println(coloresSlice)
	// Eliminar dato
	coloresSlice = append(coloresSlice[:2], coloresSlice[3:]...)
	fmt.Println(coloresSlice)
}
