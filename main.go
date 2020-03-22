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

func Draw(cards []deck.Card) (deck.Card, []deck.Card) {
  return cards[0], cards[1:]
}

func main() {
  cards := deck.New(deck.Deck(3), deck.Shuffle)
  var card deck.Card
  // for i := 0; i < 10; i++ {
  //   card, cards = cards[0], cards[1:] // Draw card, remove from deck
  //   fmt.Println(card)
  // }
  // var h Hand = cards[0:3]
  // fmt.Println(h)
  var player, dealer Hand
  for i := 0; i < 2 ; i++ { // Draw two cards
    for _, hand := range []*Hand{&player, &dealer} {
      card, cards = Draw(cards)
      *hand = append(*hand, card)
    }
  }
  fmt.Println("Player's hand:", player)
  fmt.Println("Dealer's hand:", dealer)
}
