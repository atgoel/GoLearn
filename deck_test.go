package main

import "testing"

func TestNewDeck(t *testing.T) {

	testDeck := newDeck()

	if len(testDeck) != 520 {
		t.Errorf("Expected Value 52 , but received %v", len(testDeck))
	}

}

func TestSaveToFile(t *testing.T) {

}
