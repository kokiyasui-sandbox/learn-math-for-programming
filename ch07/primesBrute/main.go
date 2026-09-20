package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: main <int>\n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Args must be integer\n")
	}

	primes := primesBrute(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}

func primesBrute(n int) []int {
	var primes []int

	for i := 2; i <= n; i++ {

		isPrime := true

		for p := 2; p < i; p++ {
			if i%p == 0 {
				isPrime = false
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	return primes
}
