package main

import "fmt"

func main() {
	amount := 1972.78
	fmt.Println(denominationChange(amount))
}

// denominationChange
func denominationChange(amount float64) error {
	// cashUnits is a slice that holds all the various cash units provided.
	// cashUnits is stored as a slice of float64.
	cashUnits := []float64{
		500.00, 200.00, 100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00,
		0.50, 0.20, 0.10, 0.05, 0.02, 0.01,
	}

}
