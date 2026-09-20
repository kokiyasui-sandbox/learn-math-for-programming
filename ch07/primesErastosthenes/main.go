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
		os.Exit(1)
	}

	if n < 2 {
		fmt.Fprintf(os.Stderr, "input must be integer\n")
		os.Exit(1)
	}

	primes := primesErastothenes(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}

func primesErastothenes(n int) []int {
	isPrimes := make([]bool, n+1)

	for i := range isPrimes {
		isPrimes[i] = true
	}

	p := 2
	for p*p <= n {
		if isPrimes[p] {
			for i := p; i <= n/p; i++ {
				isPrimes[p*i] = false
			}
		}
		p++
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrimes[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
