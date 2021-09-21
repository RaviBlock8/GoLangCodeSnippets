package main

import "testing"

func TestNewDeck(t *testing.T){
	deck:=newDeck()
	if len(deck)!=12{
		t.Errorf("Size of the created deck is not 12")
	}
}