package main

import (
	"fmt"
	// "reflect"
)

func main() {

	exp := map[string]float32{"Month pay: ": 0, "Period credit (years): ": 0, "Sum on hand: ": 0}

	for key, val := range exp {
		fmt.Print(key)
		fmt.Scan(&val)
		exp[key] = val
	}

	totalValue := exp["Period credit (years): "] * 12 * exp["Month pay: "]
	nominalValuePerc := (totalValue/exp["Sum on hand: "] - 1) * 100
	effective := nominalValuePerc / exp["Period credit (years): "]

	fmt.Printf("--------------------------\nEffective rate: %.2f\n\n", effective)

}
