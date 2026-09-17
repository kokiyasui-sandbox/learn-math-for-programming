package main

import (
	"fmt"
	"time"
)

func main() {
	seed := int(time.Now().UnixMicro() % 10_000)
	fmt.Println("seed:", seed)
	fmt.Println("random", minstd(seed))
}

func minstd(x int) float64 {
	a := 48271
	c := 2147483647

	x = x * a % c

	return float64(x) / float64(c)
}
