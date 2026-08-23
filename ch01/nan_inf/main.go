package main

import (
	"fmt"
	"math"
)

func main() {
	nanInf()
}

func nanInf() {
	fmt.Println(math.Log(-1))
	fmt.Println(math.Exp(1000))
}
