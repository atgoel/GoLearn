package main

import "fmt"

func main() {

	mydeck := newDeck()
	fmt.Println("Initialized Deck")
	mydeck.print()

	mydeck.shuffle()

	fmt.Println("First hand")
	a, _ := hand(mydeck, 10)
	a.print()

	fmt.Println("Write To file")
	mydeck.saveToFile("myFile")

}
