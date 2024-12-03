package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go TODO

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Println("Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Println("Error creating dice.log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice was rolled at", when)

}
