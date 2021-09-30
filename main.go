package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Check to see that the file name argument is passed in.
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "a floating point number is required as input")
		os.Exit(-1)
	}

	// Convert input string to float64.
	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// Pass the userInput over to the denominationChange function.
	m := denominationChange(amount)
	if m == nil {
		fmt.Fprintln(os.Stderr, "Unable to compute cashUnit")
	}

	// Range over the returned struct and format string printing out value.
	for k, v := range m {
		fmt.Fprintf(os.Stderr, "\t%d x %.2f\t= %.2f\n", v, k, float64(v)*k)
	}
}

// denominationChange returns the number of each cash-unit
// that fits into the given input by achieving the smallest total number of cash-units.
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
