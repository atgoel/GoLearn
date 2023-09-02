package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardSuit := []string{"Spade", "Club", "Hearts", "Diamonds"}
	cardRank := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			cards = append(cards, cardRank[j]+" of "+cardSuit[i])
		}
	}
	return cards

}

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func hand(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
