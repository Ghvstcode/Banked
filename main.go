package main

import (
	"fmt"
	"math"
)

func main() {
	amount := 1972.78
	fmt.Println(denominationChange(amount))
}

// denominationChange returns the number of each cash-unit
//that fits into the given input by achieving the smallest total number of cash-units.
func denominationChange(amount float64) map[float64]int {
	m := make(map[float64]int)

	// cashUnits is a slice that holds all the various cash units provided.
	// cashUnits is stored as a slice of float64.
	cashUnits := []float64{
		500.00, 200.00, 100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00,
		0.50, 0.20, 0.10, 0.05, 0.02, 0.01,
	}

	//Range over each cashUnit
	for _, val := range cashUnits {
		// if the cash unit is greater than the amount,
		//we want to skip to a lesser cash unit
		if val > amount {
			continue
		}
		// Get the integer that sums up to the amount/val
		// using the Modf function provided by the Math package
		integer, _ := math.Modf(amount / val)
		m[val] = int(integer)
		// Retrieve the remainder of the amount/cashUnit.
		amount = math.Mod(amount, val)
		if amount == 0 {
			break
		}
	}
	return m
}
