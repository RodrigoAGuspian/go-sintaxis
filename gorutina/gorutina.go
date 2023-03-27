package gorutina

import (
	"fmt"
)

func suma(nums []int, numero int, c chan int) {
	suma := 0
	for _, n := range nums {
		suma += n
	}
	fmt.Println("En", numero, "se sumó:", suma)
	c <- suma
}

func SumarTodo() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	go suma(nums[:len(nums)/2], 1, c)
	go suma(nums[len(nums)/2:], 2, c)

	x, y := <-c, <-c

	fmt.Printf("La suma total es: %d\n", x+y)
}
