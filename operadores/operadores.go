package operadores

import (
	"fmt"
	"math"
)

func OperadoresAritmeticos() {
	x := 10
	y := 2

	// Suma
	fmt.Println(x, "+", y, "=", x+y)

	// Resta
	fmt.Println(x, "-", y, "=", x-y)

	// Multipicación
	fmt.Println(x, "*", y, "=", x*y)

	// División
	fmt.Println(x, "/", y, "=", x/y)

	// Módulo
	fmt.Println(x, "%", y, "=", x%y)

	// Potencia
	fmt.Println(x, "^", y, "=", math.Pow(2, 3))

	// Incremento
	y++
	fmt.Println("y++ =", y)

	// Decremento
	x--
	fmt.Println("x-- =", x)

	// Operadores unarios
	// +
	fmt.Println("+x = ", +x)
	// -
	fmt.Println("-x = ", -x)
}

func OperadoresRelacionales() {
	x := 2
	y := 4

	// (==) Igual que
	println("x == y", x == y)
	println("y == 4", y == 4)

	// (==) Diferente que
	println("x != y", x != y)
	println("y != 4", y != 4)

	// (>) Mayor que
	println("x > y", x > y)
	println("y > x", y > x)

	// (<) Menor que
	println("x < y", x < y)
	println("y < x", y < x)

	// (>=) Mayor o igual que
	println("x >= y", x >= y)
	println("y >= x", y >= x)

	// (<=) Mayor o igual que
	println("x >= y", x <= y)
	println("y >= x", y <= x)

}

func OperadoresLogicos() {
	x := 4
	y := 2

	// AND (y)
	println("x > 2 && y < 1", x > 2 && y < 1)
	println("x == 2 && y < 1", x > 2 && y < 1)

	// OR (o)
	println("x == 2 || y < 1", x == 2 || y < 1)
	println("x == 2 || y == 1", x > 2 || y == 1)

	// NOT (Negación)
	v := true
	println("!v", !v)
	println("!!v", !!v)

}

func OperadoresLogicosBits() {
	a := byte(0b00000001) // 1
	b := byte(0b00000010) // 2

	// a y b en binario y entero
	fmt.Printf("a = %08b ; %02x\n", a, a)
	fmt.Printf("b = %08b ; %02x\n", b, b)

	// AND
	fmt.Printf("a & b = %08b ; %02x\n", a&b, a&b)

	// OR
	fmt.Printf("a | b = %08b ; %02x\n", a|b, a|b)

	// XOR
	fmt.Printf("a ^ b = %08b ; %02x\n", a^b, a^b)

	// Correr bits a la izquierda
	fmt.Printf("a << 2 = %08b ; %02x\n", a<<2, a<<2)

	// Correr bits a la derecha
	fmt.Printf("b >> 1 = %08b ; %02x\n", b>>1, b>>1)

}

func OperadoresAsignacion() {
	// (=) Asignación simple (=)
	var i int = 1
	fmt.Println("i = ", i)

	// (+=) Suma y Asignación
	i += 14
	fmt.Println("i += 14 es", i)

	// (-=) Resta y Asignación
	i -= 10
	fmt.Println("i -= 10 es", i)

	// (*=) Multiplicación y Asignación
	i *= 4
	fmt.Println("i *= 4 es", i)

	// (/=) División y Asignación
	i /= 5
	fmt.Println("i /= 5 es", i)

	// (/=) Modulo y Asignación
	i %= 5
	fmt.Println("i %= 5 es", i)

	a := byte(0b00000001) // 1
	b := byte(0b00000010) // 2

	// a y b en binario y entero
	fmt.Printf("a = %08b ; %02x\n", a, a)
	fmt.Printf("b = %08b ; %02x\n", b, b)

	// (&=) AND y Asignación
	a &= b
	fmt.Printf("a &= b --> a = %08b ; %02x\n", a, a)

	a = byte(0b00000001)

	// (|=) OR y Asignación
	a |= b
	fmt.Printf("a |= b --> a = %08b ; %02x\n", a, a)

	a = byte(0b00000001)

	// (^=) OR y Asignación
	a ^= b
	fmt.Printf("a ^= b --> a = %08b ; %02x\n", a, a)

	// (<<=) Correr bits a la izquierda y Asignación
	a <<= 2
	fmt.Printf("a <<= 2 --> a = %08b ; %02x\n", a, a)

	// (>>=) Correr bits a la derecha y Asignación
	b >>= 1
	fmt.Printf("b >= 1 --> a = %08b ; %02x\n", a, a)
}

func OperadoresDireccion() {
	a := 10
	fmt.Println("Valor de a:", a)

	// (&) Dirección de memoria
	b := &a
	fmt.Println("Dirección de a:", b)

	// (*) Valor apuntador
	c := *b
	fmt.Println("Valor donde apunta b:", c)
}
