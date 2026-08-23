package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var val float32 = 2.718

	binary32(val)
}

func binary32(val float32) {
	ptr := unsafe.Pointer(&val)
	bits := *(*uint32)(ptr)

	fmt.Printf("%1.8f\n", val)
	fmt.Printf("%032b\n", bits)
}
