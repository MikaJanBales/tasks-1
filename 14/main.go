package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int)
	arr := []any{"hi", 42, true, ch}
	for _, v := range arr {
		switch x := v.(type) {
		case int:
			fmt.Println("int:", x)
		case string:
			fmt.Println("string:", x)
		case bool:
			fmt.Println("bool:", x)
		case chan int:
			fmt.Println("channel for integer:", x)
		default:
			fmt.Println(reflect.TypeOf(v).String())
		}
	}
}
