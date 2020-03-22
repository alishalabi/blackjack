package main

import (
  "github.com/alishalabi/deck"
  "strings"
  "fmt"
)

type Hand []deck.Card

// String returns all cards in hand, in print-friendly format
func(h Hand) String() string {
  strgs := make([]string, len(h))
  for i, _ := range h {
    strgs[i] = h[i].String()
  }
  return strings.Join(strgs, ", ")
}

func main() {
  // fmt.println("Hello world")
  cards := deck.New(deck.Deck(3), deck.Shuffle)
  var card deck.Card
  for i := 0; i < 10; i++ {
    card, cards = cards[0], cards[1:] // Draw card, remove from deck
    fmt.Println(card)
  }
  var h Hand = cards[0:3]
  fmt.Println(h)
}
