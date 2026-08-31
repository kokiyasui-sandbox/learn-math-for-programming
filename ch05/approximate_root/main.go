package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: main <i>")
		os.Exit(1)
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("argument must be integer")
		os.Exit(1)
	}

	i := 0
	for num > 0 {
		n := 2*i + 1
		num = num - n

		i++
	}

	fmt.Printf("approximate root: %d\n", i)
}
