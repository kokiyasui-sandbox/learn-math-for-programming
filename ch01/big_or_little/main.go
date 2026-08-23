package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint64 = 0x123456789
	dumpMemory(a)
}

func dumpMemory(val uint64) {
	ptr := &val
	size := unsafe.Sizeof(val)

	bytes := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
	addr := uintptr(unsafe.Pointer(&val))

	for i := range size {
		fmt.Printf("0x%012x | 0x%02x \n", addr+i, bytes[i])
	}
}
