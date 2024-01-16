package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func binarySearch[T constraints.Ordered](list []T, item T, left int, right int) int {
	if left > right {
		return -1
	}

	piv := (right + left) / 2

	if item == list[piv] {
		return piv
	} else if item > list[piv] {
		return binarySearch(list, item, piv+1, right)
	} else {
		return binarySearch(list, item, left, piv-1)
	}
}

func main() {
	myList := []int{1, 3, 5, 7, 9, 13, 15, 17, 19, 20, 25, 34, 44, 45, 55, 56, 67, 78, 90, 100}
	myFloatList := []float32{1.1, 1.6, 2.5, 4.5, 6.1, 7.6, 11.4, 15.5}
	myStringList := []string{"ab", "ac", "b", "bcd", "fav", "fbd"}

	fmt.Println(binarySearch(myList, 1, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 45, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 28, 0, len(myList)-1))

	fmt.Println(binarySearch(myFloatList, 1.6, 0, len(myFloatList)-1))
	fmt.Println(binarySearch(myFloatList, 10, 0, len(myFloatList)-1))
	fmt.Println(binarySearch(myFloatList, 15.5, 0, len(myFloatList)-1))

	fmt.Println(binarySearch(myStringList, "arm", 0, len(myStringList)-1))
	fmt.Println(binarySearch(myStringList, "ab", 0, len(myStringList)-1))
	fmt.Println(binarySearch(myStringList, "fav", 0, len(myStringList)-1))

}
