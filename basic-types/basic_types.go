package main

import "fmt"

func main() {
	var a int = -42
	var b uint = 42
	var c int8 = -128
	var d uint8 = 255

	var e float32 = 3.14
	var f float64 = 3.1415926535

	var g complex64 = complex(1, 2)
	var h complex128 = complex(3, 4)

	var i byte = 'A'
	var j string = "Hello, Go!"

	var k bool = true

	fmt.Println("Integers:", a, b, c, d)
	fmt.Println("Float:", e, f)
	fmt.Println("Complex:", g, h)
	fmt.Println("Byte & string:", i, j)
	fmt.Println("Boolean:", k)
}
