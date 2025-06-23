package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Sizes of types in bytes:")
	fmt.Println("int:", unsafe.Sizeof(int(0)))
	fmt.Println("int8:", unsafe.Sizeof(int8(0)))
	fmt.Println("int16:", unsafe.Sizeof(int16(0)))
	fmt.Println("int32:", unsafe.Sizeof(int32(0)))
	fmt.Println("int64:", unsafe.Sizeof(int64(0)))
	fmt.Println("uint:", unsafe.Sizeof(uint(0)))
	fmt.Println("uint8:", unsafe.Sizeof(uint8(0)))
	fmt.Println("uint16:", unsafe.Sizeof(uint16(0)))
	fmt.Println("uint32:", unsafe.Sizeof(uint32(0)))
	fmt.Println("uint64:", unsafe.Sizeof(uint64(0)))
	fmt.Println("float32:", unsafe.Sizeof(float32(0)))
	fmt.Println("float64:", unsafe.Sizeof(float64(0)))
	fmt.Println("complex64:", unsafe.Sizeof(complex64(0)))
	fmt.Println("complex128:", unsafe.Sizeof(complex128(0)))
	fmt.Println("byte:", unsafe.Sizeof(byte(0)))
	fmt.Println("rune:", unsafe.Sizeof(rune(0)))
	fmt.Println("bool:", unsafe.Sizeof(bool(true)))
	fmt.Println("string (empty):", unsafe.Sizeof(string("")))
}
