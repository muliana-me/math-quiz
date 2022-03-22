package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	fmt.Println("\nMath Multiplication Test")
	fmt.Println("````````````````````````")

	score := 0
	questions := 10

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < questions; i++ {
		num1 := randInt(1, 10)
		num2 := randInt(1, 10)
		fmt.Printf("\n%v x %v = ", num1, num2)
		var answer uint
		fmt.Scan(&answer)
		result := num1 * num2

		if answer == uint(result) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Incorrect! The answer is %v.\n", result)
		}
	}

	fmt.Printf("\nCorrect answer %v out of %v.", score, questions)
	percent := (float64(score) / float64(questions)) * 100
	fmt.Printf("\nYour score is %v.\n\n", percent)
}
