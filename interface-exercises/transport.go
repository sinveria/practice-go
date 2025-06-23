package main

import "fmt"

type Transport interface {
	Move(destination string)
	Stop()
	GetInfo() string
}

type Car struct {
	Model string
	Speed int
}

func (c Car) Move(destination string) {
	fmt.Printf("Car %s is traveling at %s at %d km/h\n", c.Model, destination, c.Speed)
}

func (c Car) Stop() {
	fmt.Printf("Car %s has stopped\n", c.Model)
}

func (c Car) GetInfo() string {
	return fmt.Sprintf("Vehicle %s (max. speed %d km/h)", c.Model, c.Speed)
}

type Bicycle struct {
	Type string
}

func (b Bicycle) Move(destination string) {
	fmt.Printf("Bicycle (%s) is going to %s\n", b.Type, destination)
}

func (b Bicycle) Stop() {
	fmt.Printf("Bicycle (%s) has stopped\n", b.Type)
}

func (b Bicycle) GetInfo() string {
	return fmt.Sprintf("Bicycle (%s)", b.Type)
}

func UseTransport(t Transport, destination string) {
	fmt.Println("\nVehicle:", t.GetInfo())
	t.Move(destination)
	t.Stop()
}

func main() {
	car := Car{Model: "Toyota Camry", Speed: 120}
	bike := Bicycle{Type: "mountain"}

	UseTransport(car, "Moscow")
	UseTransport(bike, "park")

	transports := []Transport{car, bike}
	fmt.Println("\nAll vehicles:")
	for _, t := range transports {
		UseTransport(t, "destination")
	}
}
