package main

import "fmt"

func main() {

	var firstName string = "Pasacl"
	var lastName string = "Fratamico"

	var dayOfBirth int = 18
	var monthOfBirth int = 1
	var yearOfBirth int = 2005
	var numberOfSiblings int = 3
	var heightInMeters float32 = 1.8
	var zodiacSign = '\u2651'

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
