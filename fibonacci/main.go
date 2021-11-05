package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// play `go run . 10`
func main() {
	log.SetPrefix("fibonacci: ")
	log.SetFlags(0)

	if len(os.Args) == 1 {
		log.Fatal(errors.New("arguments length is equal or over 2"))
		os.Exit(1)
	}
	fibonacciNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if fibonacciNumber == 0 {
		log.Fatal(errors.New("args[1] is equal or over 1"))
		os.Exit(1)
	}

	fib := fibonacci(fibonacciNumber)
	log.Print(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(fib)), ", "), "[]"))
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
