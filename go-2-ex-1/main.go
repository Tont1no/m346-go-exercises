package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			FirstName: "Pascal",
			LastName:  "Fratamico",
		},
		BirthDate: BirthDate{
			Day:   18,
			Month: 1,
			Year:  2005,
		},
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u2651', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
