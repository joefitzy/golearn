// Function Basics
package main

import (
	"errors"
	"fmt"
)

// optional init function
// runs after package is loaded and before main runs
// can have multiple inits
// and order matters in code.
// go for setting package state
func init() {
	fmt.Println("Veni.")
}

func init() {
	fmt.Println("Vidi.")
}

func init() {
	fmt.Println("Vici.")
}

func printMe(message string) {
	fmt.Println(message)
}

func main() {
	message := "Hello World!"
	printMe(message)

	var result, remainder, err = intDivision(24, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result of integer division is %v", result)
	} else {
		fmt.Printf("Result of integer division is %v with %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by Zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
