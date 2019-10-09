package main

import (
	"math/rand"
	"time"
)

// Card represents all possible cards in Too Many Monkees
type Card uint8

// Flip will reveal a hidden card, or vice versa
func (c Card) Flip() Card {
	return c ^ CardHidden
}

// Hide will mark a shown card as hidden. No-op if already hidden
func (c Card) Hide() Card {
	return c | CardHidden
}

// Show will mark a hidden card as shown. No-op if already showing
func (c Card) Show() Card {
	return c & CardVisible
}

// IsNumber indicates that the card is a number card
func (c Card) IsNumber() bool {
	return c <= CardSix
}

// IsHidden indicates that the player cannot tell the meaning of the card
func (c Card) IsHidden() bool {
	return c&CardHidden == CardHidden
}

// Deck is a list of cards
type Deck []Card

type cardCount struct {
	card  Card
	count int
}

const (
	// CardOne is a card with one monkey
	CardOne = 0
	// CardTwo is a card with two monkeys
	CardTwo = 1
	// CardThree is a card with three monkeys
	CardThree = 2
	// CardFour is a card with four monkeys
	CardFour = 3
	// CardFive is a card with five monkeys
	CardFive = 4
	// CardSix is a card with six monkeys
	CardSix = 5
	// CardWild is a card that can stand in for any number card
	CardWild = 'W'
	// CardSkip is a card that can be used to skip an opponent's turn
	CardSkip = 'S'
	// CardRacoon allows the player to take any card from the discard pile
	CardRacoon = 'R'
	// CardGiraffe can't come to the party
	CardGiraffe = 'G'
	// CardElephant can't come to the party
	CardElephant = 'E'
	// CardDoNotDisturb hides one of the player's showing cards
	CardDoNotDisturb = 'D'
	// CardHidden is a flag on any card indicating that it is visible
	CardHidden = 128
	// CardVisible is a bitmask that can be used to show a card
	CardVisible = 127
	// CardOneU is CardOne, but hidden
	CardOneU = CardOne | CardHidden
	// CardTwoU is CardTwo, but hidden
	CardTwoU = CardTwo | CardHidden
	// CardThreeU is CardThree, but hidden
	CardThreeU = CardThree | CardHidden
	// CardFourU is CardFour, but hidden
	CardFourU = CardFour | CardHidden
	// CardFiveU is CardFive, but hidden
	CardFiveU = CardFive | CardHidden
	// CardSixU is CardSix, but hidden
	CardSixU = CardSix | CardHidden
	// CardWildU is CardWild, but hidden
	CardWildU = CardWild | CardHidden
	// CardSkipU is CardSkip, but hidden
	CardSkipU = CardSkip | CardHidden
	// CardRacoonU is CardRacoon, but hidden
	CardRacoonU = CardRacoon | CardHidden
	// CardGiraffeU is CardGiraffe, but hidden
	CardGiraffeU = CardGiraffe | CardHidden
	// CardElephantU is CardElephant, but hidden
	CardElephantU = CardElephant | CardHidden
	// CardDoNotDisturbU is CardDoNotDisturb, but hidden
	CardDoNotDisturbU = CardDoNotDisturb | CardHidden
	// CardNone is a lack of card -- the null card
	CardNone = 255
)

// String converts any card to its string representation
func (c Card) String() string {
	switch c {
	case CardOne:
		return "1"
	case CardTwo:
		return "2"
	case CardThree:
		return "3"
	case CardFour:
		return "4"
	case CardFive:
		return "5"
	case CardSix:
		return "6"
	case CardWild:
		return "W"
	case CardSkip:
		return "S"
	case CardRacoon:
		return "R"
	case CardGiraffe:
		return "G"
	case CardElephant:
		return "E"
	case CardDoNotDisturb:
		return "D"

	case CardOneU:
		return "u1"
	case CardTwoU:
		return "u2"
	case CardThreeU:
		return "u3"
	case CardFourU:
		return "u4"
	case CardFiveU:
		return "u5"
	case CardSixU:
		return "u6"
	case CardWildU:
		return "uW"
	case CardSkipU:
		return "uS"
	case CardRacoonU:
		return "uR"
	case CardGiraffeU:
		return "uG"
	case CardElephantU:
		return "uE"
	case CardDoNotDisturbU:
		return "uD"
	}
	panic("unknown card")
}

var deckCounts = []cardCount{
	cardCount{CardOne, 6},
	cardCount{CardTwo, 6},
	cardCount{CardThree, 6},
	cardCount{CardFour, 6},
	cardCount{CardFive, 6},
	cardCount{CardSix, 6},
	cardCount{CardWild, 4},
	cardCount{CardGiraffe, 3},
	cardCount{CardElephant, 3},
	cardCount{CardSkip, 4},
	cardCount{CardDoNotDisturb, 2},
	cardCount{CardRacoon, 3},
}

var sizeOfDeck = sizeOfDeckCalc()

func sizeOfDeckCalc() int {
	cnt := 0
	for _, cc := range deckCounts {
		cnt += cc.count
	}
	return cnt
}

// NewDeck creates a deck of all the cards in a game of Too Many Monkeys
func NewDeck() Deck {
	deck := make(Deck, sizeOfDeck)
	ix := 0
	for _, cc := range deckCounts {
		for i := 0; i < cc.count; i++ {
			deck[ix] = cc.card
			ix++
		}
	}
	return deck
}

// Shuffle randomizes the order of a Deck
func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(d); i++ {
		target := i + r.Intn(len(d)-i)
		d[i], d[target] = d[target], d[i]
	}
}
