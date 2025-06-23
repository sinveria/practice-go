package main

import "fmt"

func incrementByValue(a int) {
	a++
	fmt.Println("Inside incrementByValue:", a)
}

func incrementByPointer(a *int) {
	*a++
	fmt.Println("Inside incrementByPointer:", *a)
}

type Person struct {
	Name string
	Age  int
}

func main() {
	num := 10

	fmt.Println("Original value:", num)

	incrementByValue(num)
	fmt.Println("After incrementByValue:", num)

	incrementByPointer(&num)
	fmt.Println("After incrementByPointer:", num)

	p := Person{"Alexander", 25}

	fmt.Println("\nOriginal person:", p)

	modifyPersonByValue(p)
	fmt.Println("After modifyPersonByValue:", p)

	modifyPersonByPointer(&p)
	fmt.Println("After modifyPersonByPointer:", p)
}

func modifyPersonByValue(p Person) {
	p.Age += 5
	fmt.Println("Inside modifyPersonByValue:", p)
}

func modifyPersonByPointer(p *Person) {
	p.Age += 5
	fmt.Println("Inside modifyPersonByPointer:", *p)
}
