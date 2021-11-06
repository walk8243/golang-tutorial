package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pi(10000)) // This is accuracy
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go term(ch, float64(k)) // open other channel
	}
	f := 0.0
	for k := 0; k < n; k++ {
		f += <-ch // get other channel
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
