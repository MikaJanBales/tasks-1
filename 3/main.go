package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	var res int

	for _, el := range nums {
		wg.Add(1)
		go func(el int) {
			defer wg.Done()
			res += el * el
		}(el)
	}

	wg.Wait()

	fmt.Print(res)
}
