package main

import (
	"log"
)

func main() {
	log.SetPrefix("prime-number: ")
	log.SetFlags(0)

	primeNumbers := getPrimeNumber(100)
	log.Println(primeNumbers)
}

func getPrimeNumber(num int) []int {
	target, results := 2, []int{}

	for i := 0; i < num; i++ {
		for {
			isPrime := true
			for _, v := range results {
				if target%v == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				results = append(results, target)
			}
			target++
			if isPrime {
				break
			}
		}
	}

	return results
}
