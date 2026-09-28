package main

import "fmt"

type Vehicle struct {
	number string
	rate   float64
}

type Car struct {
	Vehicle
	hours int
}

type Bike struct {
	Vehicle
	hours int
}

func (c Car) calculate() {
	fmt.Println("Smart Parking Manager")
	fmt.Println("\n--- Car Parking ---")
	fmt.Println("Vehicle No:", c.number)
	fmt.Println("Hours:", c.hours)
	fmt.Printf("Parking Fee: Rs. %.0f\n", c.rate*float64(c.hours))
}

func (b Bike) calculate() {
	fmt.Println("Smart Parking Manager")
	fmt.Println("\n--- Bike Parking ---")
	fmt.Println("Vehicle No:", b.number)
	fmt.Println("Hours:", b.hours)
	fmt.Printf("Parking Fee: Rs. %.0f\n", b.rate*float64(b.hours))
}

func main() {
	car := Car{
		Vehicle: Vehicle{"GJ01AB1234", 30},
		hours:   4,
	}

	bike := Bike{
		Vehicle: Vehicle{"GJ01XY5678", 15},
		hours:   3,
	}

	car.calculate()
	bike.calculate()
}
