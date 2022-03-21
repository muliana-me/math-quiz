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
	fmt.Println("\nPerkalian Matematika")
	fmt.Println("````````````````````")

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
			fmt.Println("Benar!")
			score++
		} else {
			fmt.Printf("Salah! Jawaban yang benar %v.\n", result)
		}
	}

	fmt.Printf("\nTotal jawaban benar %v dari %v soal.", score, questions)
	percent := (float64(score) / float64(questions)) * 100
	fmt.Printf("\nMendapatkan nilai %v.\n\n", percent)
}
