package main

import (
  "github.com/alishalabi/deck"
  "strings"
  "fmt"
)

type Hand []deck.Card

// String returns all cards in hand, in print-friendly format
func (h Hand) String() string {
  strgs := make([]string, len(h))
  for i, _ := range h {
    strgs[i] = h[i].String()
  }
  return strings.Join(strgs, ", ")
}

// DealerString hides the dealer's second hand
func (h Hand) DealerString() string {
  return h[0].String() + ", **Hidden**"
}

// Score returns score, with Aces having valiable values
func (h Hand) Score() int {
  minScore := h.MinScore()
  if minscore > 11 { // 10 is the highest time we would want Ace == 11
    return minScore
  }
  for _, c := range h {
    if c.Rank == deck.Ace {
      return minScore + 10 // Ace is already worth 1, adding 10 makes it work 10
    }
  }
}

// MinScore calculates the minimum score for a hand (ie Ace equals 1)
func (h Hand) MinScore() int {
  score := 0
  for _, c := range h {
    score += min(int(c.Rank), 10) // Any rank higher than 10 will return 10
  }
  return score
}

// min is an internal helper function that helps convert Jack, Queen, King to value ten
func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

// Draw takes the top card of the deck, then returns that card and a diminshed deck
func draw(cards []deck.Card) (deck.Card, []deck.Card) {
  return cards[0], cards[1:]
}

func main() {
  cards := deck.New(deck.Deck(3), deck.Shuffle)
  var card deck.Card
  var player, dealer Hand
  for i := 0; i < 2 ; i++ { // Draw two cards
    for _, hand := range []*Hand{&player, &dealer} { // Using pointers to avoid "hand" copies
      card, cards = draw(cards)
      *hand = append(*hand, card)
    }
  }

  // Game run logic
  var input string
  for input != "s" {
    fmt.Println("Player's hand:", player)
    fmt.Println("Dealer's hand:", dealer.DealerString())
    fmt.Println("Player's choice: (h)it or (s)tand")
    fmt.Scanf("%s\n", &input)
    switch input {
    case "h":
      card, cards = draw(cards)
      player = append(player, card)
    default:
      fmt.Println("Whoops, not a valid option. Please enter *h* to hit, or *s* to stand.")
    }
  }
  fmt.Println("***Final Hands***")
  fmt.Println("Player's final hand:", player, "\nScore:", player.Score())
  fmt.Println("Dealer's final hand:", dealer, "\nScore:", dealer.Score())


}
