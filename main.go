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
  if minScore > 11 { // 10 is the highest time we would want Ace == 11
    return minScore
  }
  for _, c := range h {
    if c.Rank == deck.Ace {
      return minScore + 10 // Ace is already worth 1, adding 10 makes it work 10
    }
  }
  return minScore
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

type State int8

const (
  StatePlayerTurn State = iota
  StateDealerTurn
  StateHandOver
)

type GameState struct {
  Deck []deck.Card
  State State
  Player Hand
  Dealer Hand
}

// CurrentPlayer returns pointer to current player's hand
func (gs *Gamestate) CurrentPlayer() *Hand {
  switch gs.State {
  case StatePlayerTurn:
    return &gs.Player
  case StateDealerTurn:
    return &gs.Dealer
  default:
    panic("No valid player")
  }
}


// clone returns a shallow copy of a GameState
func clone(gs GameState) GameState {
  ret := GameState {
    Deck: make([]deck.Card, len(gs.Deck)),
    State: gs.State,
    Player: make(Hand, len(gs.Player)),
    Dealer: make(Hand, len(gs.Dealer)),
  }
copy(ret.Deck, gs.Deck)
copy(ret.Player, gs.Player)
copy(re.Dealer, gs.Dealer)
return ret
}

func main() {
  var gs GameState
  gs.Deck = deck.New(deck.Deck(3), deck.Shuffle)

  // cards := deck.New(deck.Deck(3), deck.Shuffle)
  // var card deck.Card
  // var player, dealer Hand
  // for i := 0; i < 2 ; i++ { // Draw two cards
  //   for _, hand := range []*Hand{&player, &dealer} { // Using pointers to avoid "hand" copies
  //     card, cards = draw(cards)
  //     *hand = append(*hand, card)
  //   }
  // }
  //
  // // Game run logic
  // var input string
  // for input != "s" {
  //   fmt.Println("Player's hand:", player)
  //   fmt.Println("Dealer's hand:", dealer.DealerString())
  //   fmt.Println("Player's choice: (h)it or (s)tand")
  //   fmt.Scanf("%s\n", &input)
  //   switch input {
  //   case "h":
  //     card, cards = draw(cards)
  //     player = append(player, card)
  //   // default:
  //   //   fmt.Println("Whoops, not a valid option. Please enter *h* to hit, or *s* to stand.")
  //   }
  // }
  // // Basic dealer AI
  // // Hit if dealer has less than 16, or a soft 17 (17 with high Ace)
  // for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
  //   card, cards = draw(cards)
  //   dealer = append(dealer, card)
  //   fmt.Println("Dealer draws:", card)
  // }
  // pScore, dScore := player.Score(), dealer.Score()
  // fmt.Println("\n***Final Hands***")
  // fmt.Println("Player's final hand:", player, "\nScore:", pScore)
  // fmt.Println("Dealer's final hand:", dealer, "\nScore:", dScore)
  // switch { // Calculate results
  // case pScore > 21:
  //   fmt.Println("Player busts")
  // case dScore > 21:
  //   fmt.Println("Dealer busts")
  // case pScore > dScore:
  //   fmt.Println("Player wins")
  // case dScore > pScore:
  //   fmt.Println("Dealer wins")
  // case pScore == dScore:
  //   fmt.Println("Player and Dealer draw")
  // }

}
