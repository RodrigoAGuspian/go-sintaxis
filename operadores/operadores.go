package operadores

import (
	"fmt"
	"math"
)

func OperadoresAritmeticos() {
	x := 10
	y := 2

	fmt.Println(x, "+", y, "=", x+y) // Suma
	fmt.Println(x, "-", y, "=", x-y) // Resta
	fmt.Println(x, "*", y, "=", x*y) // Multipicación
	fmt.Println(x, "/", y, "=", x/y) // División
	fmt.Println(x, "%", y, "=", x%y) // Módulo
	fmt.Println(x, "^", y, "=",
		math.Pow(float64(x), float64(y))) // Potencia

	y++ // Incremento
	fmt.Println("y++ =", y)
	x-- // Decremento
	fmt.Println("x-- =", x)

	// Operadores unarios
	fmt.Println("+x = ", +x) // +
	fmt.Println("-x = ", -x) // -
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
	// (<=) Menor o igual que
	println("x <= y", x <= y)
	println("y <= x", y <= x)

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

	fmt.Printf("a & b = %08b ; %02x\n", a&b, a&b)    // AND
	fmt.Printf("a | b = %08b ; %02x\n", a|b, a|b)    // OR
	fmt.Printf("a ^ b = %08b ; %02x\n", a^b, a^b)    // XOR
	fmt.Printf("a << 2 = %08b ; %02x\n", a<<2, a<<2) // Correr bits a la izquierda
	fmt.Printf("b >> 1 = %08b ; %02x\n", b>>1, b>>1) // Correr bits a la derecha

}

func OperadoresAsignacion() {
	var i int = 1 // (=) Asignación simple (=)
	var tmp int = 0
	fmt.Println("i = ", i)

	tmp = i
	i += 14 // (+=) Suma y Asignación
	fmt.Println(tmp, "+= 14 es", i)

	tmp = i
	i -= 10 // (-=) Resta y Asignación
	fmt.Println(tmp, "-= 10 es", i)

	tmp = i
	i *= 4 // (*=) Multiplicación y Asignación
	fmt.Println(tmp, "*= 4 es", i)

	tmp = i
	i /= 5 // (/=) División y Asignación
	fmt.Println(tmp, "/= 5 es", i)

	tmp = i
	i %= 5 // (/=) Modulo y Asignación
	fmt.Println(tmp, "%= 5 es", i)

	a := byte(0b00000001) // 1
	b := byte(0b00000010) // 2

	// a y b en binario y entero
	fmt.Printf("a = %08b ; %02x\n", a, a)
	fmt.Printf("b = %08b ; %02x\n", b, b)

	// (&=) AND y Asignación
	a &= b
	fmt.Printf("a &= b --> a = %08b ; %d\n", a, a)

	a = byte(0b00000001)

	// (|=) OR y Asignación
	a |= b
	fmt.Printf("a |= b --> a = %08b ; %d\n", a, a)

	a = byte(0b00000001)

	// (^=) XOR y Asignación
	a ^= b
	fmt.Printf("a ^= b --> a = %08b ; %d\n", a, a)

	a = byte(0b00000001)
	// (<<=) Correr bits a la izquierda y Asignación
	a <<= 2
	fmt.Printf("a <<= 2 --> a = %08b ; %d\n", a, a)

	// (>>=) Correr bits a la derecha y Asignación
	b >>= 1
	fmt.Printf("b >= 1 --> a = %08b ; %d\n", a, a)
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
