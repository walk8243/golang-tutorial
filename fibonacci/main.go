package main

import (
	"fmt"
)

func main() {
	fib := fibonacci(10)
	fmt.Println(fib)
}

func fibonacci(num int) []int {
	a, b := 0, 1
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		a, b = b, a+b
		arr[i] = a
	}

	return arr
}
