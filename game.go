package main

import (
	"math/rand"
)

// Game represents a game of Too Many Monkeys
type Game struct {
	Players  []*Player
	deck     Deck
	discards Deck
	rnd      *rand.Rand
}

// NewGame creates a game with a set of players. rnd is provided so that the
// caller can ensure that it isn't used across goroutines
func NewGame(players []*Player, rnd *rand.Rand) *Game {
	if len(players) == 0 {
		panic("At least one player is required")
	}

	deck := NewDeck()
	deck.Shuffle()
	for _, player := range players {
		for cardIx := 0; cardIx < len(player.cards); cardIx++ {
			player.cards[cardIx] = deck[0].Hide()
			deck = deck[1:]
		}
	}

	return &Game{players, deck, Deck{}, rnd}
}

// PeekDiscard returns the top card in the discard pile
func (game *Game) PeekDiscard() Card {
	if len(game.discards) == 0 {
		return CardNone
	}
	return game.discards[len(game.discards)-1]
}

// TakeDiscard removes and returns the top card in the discard pile
func (game *Game) TakeDiscard() Card {
	if len(game.discards) == 0 {
		return CardNone
	}
	lastIx := len(game.discards) - 1
	topCard := game.discards[lastIx]
	game.discards = game.discards[:lastIx]
	return topCard
}

// TakeDeck returns the top card from the set of unknown cards
func (game *Game) TakeDeck() Card {
	if len(game.deck) == 0 {
		game.deck, game.discards = game.discards, game.deck
		game.deck.Shuffle()
	}

	topCard := game.deck[0]
	game.deck = game.deck[1:]
	return topCard
}

// Discard puts a card on top of the discard pile
func (game *Game) Discard(card Card) {
	game.discards = append(game.discards, card)
}
