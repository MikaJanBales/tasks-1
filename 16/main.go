package main

import "fmt"

func quickSort(m []int, left int, right int) []int {
	if right <= left {
		return m
	}

	wall := left

	for i := left; i <= right; i++ {
		if m[right] >= m[i] {
			if m[i] < m[wall] {
				m[i], m[wall] = m[wall], m[i]
			}
			wall++
		}
	}

	quickSort(m, left, wall-2)
	quickSort(m, wall, right)

	return m
}

func main() {
	m := []int{7, 2, 5, 0, 1, 8, 3, 6, 9, 4}

	fmt.Print(quickSort(m, 0, len(m)-1))
}
