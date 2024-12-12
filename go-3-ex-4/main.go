package main

import "fmt"

const (
	Diamonds = '\u25c6' // Karo
	Spades   = '\u2660' // Pik
	Clubs    = '\u2663' // Kreuz
	Hearts   = '\u2665' // Herz

	Six   = "⑥"
	Seven = "⑦"
	Eight = "⑧"
	Nine  = "⑨"
	Ten   = "⑩"
	Jack  = "J"
	Queen = "Q"
	King  = "K"
	Ace   = "A"
)

func main() {
	suits := []rune{Diamonds, Spades, Clubs, Hearts}
	ranks := []string{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

	fmt.Println("Direkte Ausgabe ohne Map:")
	for _, rank := range ranks {
		for _, suit := range suits {
			// Ausgabe
			fmt.Printf("%c%s\t", suit, rank)
		}
		fmt.Println()
	}

	// Die Map "cards"
	cards := map[rune][]string{
		Diamonds: {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Spades:   {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Clubs:    {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
		Hearts:   {Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace},
	}

	// Ausgabe der Karten mit der Map
	fmt.Println("\nAusgabe mit der Map:")
	for i := 0; i < len(cards[Diamonds]); i++ {
		for suit, ranks := range cards {
			// Ausgabe
			card := fmt.Sprintf("%c%s", suit, ranks[i])
			fmt.Print(card + "\t")
		}
		fmt.Println()
	}
}
