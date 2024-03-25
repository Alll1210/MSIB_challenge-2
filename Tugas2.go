package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Declare variables to store weight and height for Mark and John
	var (
		bbMark, tbMark, bbJohn, tbJohn float64
	)

	// Function to get user input for weight and height
	getUserInput("Mark", &bbMark, &tbMark)
	getUserInput("John", &bbJohn, &tbJohn)

	// Calculate BMI for Mark and John
	markBMI := calculateBMI(bbMark, tbMark)
	johnBMI := calculateBMI(bbJohn, tbJohn)

	// Determine if Mark has a higher BMI than John
	markHigherBMI := markBMI > johnBMI

	// Output the BMIs and comparison result
	fmt.Println("BMI Mark adalah:", markBMI)
	fmt.Println("BMI John adalah:", johnBMI)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John?", markHigherBMI)
}

// Function to get user input for weight and height
func getUserInput(name string, weight, height *float64) {
	var input string
	var err error

	// Get weight input
	fmt.Printf("Masukkan berat badan %s (dalam kg): ", name)
	fmt.Scanln(&input)
	*weight, err = strconv.ParseFloat(input, 64)
	handleError(err)

	// Get height input
	fmt.Printf("Masukkan tinggi badan %s (dalam meters): ", name)
	fmt.Scanln(&input)
	*height, err = strconv.ParseFloat(input, 64)
	handleError(err)
}

// Function to calculate BMI
func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

// Function to handle errors
func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
