package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)

	for i, el := range nums {
		wg.Add(1)
		go func(i int, el int) {
			defer wg.Done()
			nums[i] = el * el
		}(i, el)
	}

	wg.Wait()

	fmt.Println(nums)
}
