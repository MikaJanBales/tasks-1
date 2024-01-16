package main

import (
	"fmt"
)

func deleteElem[T any](arr []T, i int) []T {
	return append(arr[:i], arr[i+1:]...)
}

func deleteElemNoOrder[T any](arr []T, i int) []T {
	arr[i] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr[:len(arr)-1]
}

func main() {
	// удаление элемента
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("      original slice: %v\n", a)
	a = deleteElem(a, 2)
	fmt.Printf("deleted 3-rd element: %v\n\n", a)

	// удаление элемента без сохранения порядка
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("      original slice: %v\n", a)
	a = deleteElemNoOrder(a, 2)
	fmt.Printf("deleted 3-rd element: %v (no order)\n\n", a)
}
