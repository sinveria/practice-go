package main

import "fmt"

type Engine struct {
	Type         string
	Power        int
	Displacement float64
}

type Car struct {
	Mark      string
	Model     string
	Year      int
	Price     float64
	CarEngine Engine
}

func (c Car) StartEngine() {
	fmt.Printf("Starting the engine %s %s\n", c.Mark, c.Model)
}

func (c Car) PrintSpecs() {
	fmt.Printf("\nMark: %s\n", c.Mark)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Year of issue: %d\n", c.Year)
	fmt.Printf("Price: $%.2f\n", c.Price)
	fmt.Println("\nEngine:")
	fmt.Printf("  Type: %s\n", c.CarEngine.Type)
	fmt.Printf("  Power: %d HP\n", c.CarEngine.Power)
	fmt.Printf("  Displacement: %.1f L\n", c.CarEngine.Displacement)
}

func main() {
	v6Engine := Engine{
		Type:         "V6",
		Power:        300,
		Displacement: 3.0,
	}

	myCar := Car{
		Mark:      "Toyota",
		Model:     "Camry",
		Year:      2022,
		Price:     35000.00,
		CarEngine: v6Engine,
	}

	myCar.PrintSpecs()
	myCar.StartEngine()

	bmw := Car{
		Mark:  "BMW",
		Model: "X5",
		Year:  2023,
		Price: 65000.00,
		CarEngine: Engine{
			Type:         "V8",
			Power:        450,
			Displacement: 4.4,
		},
	}

	bmw.PrintSpecs()
}
