package main

import (
	"fmt"
)

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) EstimateDistance() float64 {
	fuelConsumption := 1.5 // konsumsi bahan bakar dalam liter per mill
	estimatedDistance := c.fuelin / fuelConsumption
	return estimatedDistance
}

func main() {
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	estimatedDistance := car.EstimateDistance()
	fmt.Printf("Car type: %s , est. distance: %.2f\n", car.carType, estimatedDistance)
}
