package main

import (
	"fmt"
)

func binarySearch(list []int, item int, left int, right int) int {
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

	fmt.Println(binarySearch(myList, 100, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 1, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 15, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 28, 0, len(myList)-1))
	fmt.Println(binarySearch(myList, 0, 0, len(myList)-1))
}
