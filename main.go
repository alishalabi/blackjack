package main

import (
  "github.com/alishalabi/deck"
  "fmt"
)

func main() {
  // fmt.println("Hello world")
  cards := deck.New(deck.Deck(3), deck.Shuffle)
  var card deck.Card
  for i := 0; i < 10; i++ {
    card, cards = cards[0], cards[1:] // Draw card, remove from deck
    fmt.Println(card)
  }
}
