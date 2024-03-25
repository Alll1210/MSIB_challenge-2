package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Initialize scores
	var lumbaScores, koalaScores []int

	reader := bufio.NewReader(os.Stdin)

	// Input scores for both teams
	inputScores(reader, "Lumba-lumba", &lumbaScores)
	inputScores(reader, "Koala", &koalaScores)

	// Calculate total scores
	lumbaTotal := calculateTotal(lumbaScores)
	koalaTotal := calculateTotal(koalaScores)

	// Output total scores
	fmt.Printf("\nTotal Skor Tim Lumba-lumba: %d\n", lumbaTotal)
	fmt.Printf("Total Skor Tim Koala: %d\n", koalaTotal)

	// Calculate average scores
	lumbaAverage := calculateAverage(lumbaTotal)
	koalaAverage := calculateAverage(koalaTotal)

	// Output average scores
	fmt.Printf("\nSkor rata-rata tim Lumba-lumba: %.2f\n", lumbaAverage)
	fmt.Printf("Skor rata-rata tim Koala: %.2f\n", koalaAverage)

	// Determine the winner or if it's a draw
	printResult(lumbaAverage, koalaAverage)
}

// Function to input scores for a team
func inputScores(reader *bufio.Reader, team string, scores *[]int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Masukkan skor Tim %s ke-%d: ", team, i)
		score := getInput(reader)
		*scores = append(*scores, score)
	}
}

// Function to calculate total score
func calculateTotal(scores []int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}

// Function to calculate average score
func calculateAverage(total int) float64 {
	return float64(total) / 3
}

// Function to get input from user
func getInput(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value, _ := strconv.Atoi(input)
	return value
}

// Function to print the result of the match
func printResult(lumbaAvg, koalaAvg float64) {
	switch {
	case lumbaAvg > koalaAvg:
		fmt.Println("\nTim Lumba-lumba Menang!")
	case koalaAvg > lumbaAvg:
		fmt.Println("\nTim Koala Menang!")
	default:
		fmt.Println("\nPertandingan Seri")
	}
}
