package main
import {
	fmt
	strings
	
}

type deck []string

func newDeck() deck {
	cardSuit:={"Spade", "Club", "Hearts", "Diamonds"}
	cardRank := {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i:=0; i<4; i++ {
        for j:=0; j<13; j++ {
            deck = append(deck, cardRank[j]+" of "+cardSuit[i])
        }
    }
	return deck

}

func (d deck) print{
	for i, c := range deck
	{
        fmt.Println(i, c)
    }
}

func (d deck) toString string{
	return strings.join(deck,",")
}
func hand(d deck, int size) (deck, deck){
	return d[:size], d[size:]
}

func (d deck) shuffle(){


}