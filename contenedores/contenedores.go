package contenedores

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

func Listas() {
	lista := list.New()

	lista.PushBack(10)
	lista.PushFront(12)

	lista1 := list.New()
	lista1.PushFront(14)
	lista.PushBackList(lista1)

	lista2 := list.New()
	lista2.PushFront(16)
	lista.PushFrontList(lista2)

	for item := lista.Front(); item != nil; item = item.Next() {
		fmt.Println(item)
	}
}

func Anillo() {
	ring1 := ring.New(3)

	lenRing1 := ring1.Len()

	for i := 0; i < lenRing1; i++ {
		ring1.Value = "Numero " + strconv.Itoa(i)
		ring1 = ring1.Next()
	}

	for i := 0; i < lenRing1+1; i++ {
		fmt.Println(ring1.Value)
		ring1 = ring1.Next()
	}
}
