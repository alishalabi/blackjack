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
func (gs *GameState) CurrentPlayer() *Hand {
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
copy(ret.Dealer, gs.Dealer)
return ret
}

// Shuffle resets the game deck
func Shuffle(gs GameState) GameState {
  ret := clone(gs)
  ret.Deck = deck.New(deck.Deck(3), deck.Shuffle)
  return ret
}

// Deal gives player and dealer new hands
func Deal(gs GameState) GameState {
  ret := clone(gs)
  ret.Player = make(Hand, 0, 5)
  ret.Dealer = make(Hand, 0, 5)
  var card deck.Card
  for i := 0; i < 2; i++ {
    card, ret.Deck = draw(ret.Deck)
    ret.Player = append(ret.Player, card)
    card, ret.Deck = draw(ret.Deck)
    ret.Dealer = append(ret.Dealer, card)
  }
  ret.State = StatePlayerTurn
  return ret
}

// Hit gives the player or dealer one deck, removes that card from deck
func Hit(gs GameState) GameState {
  ret := clone(gs)
  hand := ret.CurrentPlayer()
  var card deck.Card
  card, ret.Deck = draw(ret.Deck)
  *hand = append(*hand, card)
  if hand.Score() > 21 {
    return Stand(ret)
  }
  return ret
}

func EndHand(gs GameState) GameState {
  ret := clone(gs)
  pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
  fmt.Println("\n***Final Hands***")
  fmt.Println("Player's final hand:", ret.Player, "\nScore:", pScore)
  fmt.Println("Dealer's final hand:", ret.Dealer, "\nScore:", dScore)
  switch { // Calculate results
  case pScore > 21:
    fmt.Println("Player busts, dealer wins :(")
  case dScore > 21:
    fmt.Println("Dealer busts, player wins!")
  case pScore > dScore:
    fmt.Println("Player wins!")
  case dScore > pScore:
    fmt.Println("Dealer wins :(")
  case pScore == dScore:
    fmt.Println("Player and Dealer draw")
  }
  // Clear both hands
  fmt.Println()
  ret.Player = nil
  ret.Dealer = nil
  return ret
}

// Stand changes player/dealer's turn
func Stand(gs GameState) GameState {
  ret := clone(gs)
  ret.State++ // Works b/c we used iota in states
  return ret
}

func main() {
  var gs GameState
  gs = Shuffle(gs)
  gs = Deal(gs)

  // // Game run logic for player
  var input string
  for gs.State == StatePlayerTurn {
      fmt.Println("Player's hand:", gs.Player)
      fmt.Println("Dealer's hand:", gs.Dealer.DealerString())
      fmt.Println("Player's choice: (h)it or (s)tand")
      fmt.Scanf("%s\n", &input)
      switch input {
      case "h":
        gs = Hit(gs)
      case "s":
        gs = Stand(gs)
      default:
        fmt.Println("Whoops, not a valid option. Please enter *h* to hit, or *s* to stand.")
      }
  }

  // Game logic for dealer
  for gs.State == StateDealerTurn {
    // Hit if dealer has less than 16, or a soft 17 (17 with high Ace)
    if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
      gs = Hit(gs)
    } else {
      gs = Stand(gs )
    }
  }
  gs = EndHand(gs)

}
