package estructurasRepetitivas

import "fmt"

func CicloFor() {
	suma := 0
	for i := 1; i < 5; i++ {
		suma += i
		fmt.Println("valor de i:", i)
	}
	fmt.Println(suma)
}

func CicloWhile() {
	n := 1

	for n < 5 {
		n *= 2
	}

	fmt.Println(n)
}

func CicloDoWhile() {
	n := 0
	for {
		n++
		if n >= 10 {
			break
		}
	}
	println(n)

}

func CicloForEach() {
	palabras := []string{"Hola", "amigos"}
	for i, s := range palabras {
		fmt.Println(i, s)
	}
}
