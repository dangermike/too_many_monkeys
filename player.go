package main

// Player represents a participant in a game of Too Many Monkeys
type Player struct {
	id      int
	cards   []Card
	hasSkip bool
}

// IsDone indicates whether all of a players cards are exposed, which is a win
func (player *Player) IsDone() bool {
	for _, card := range player.cards {
		if card.IsHidden() {
			return false
		}
	}
	return true
}

// Turn lets a player attempt to take a card from the deck (or discards) and
// apply it to their hand
func (player *Player) Turn(game *Game) {
	if player.hasSkip {
		println("I've been skipped!")
		player.hasSkip = false
		game.Discard(CardSkip)
		return
	}
	topDiscard := game.PeekDiscard()
	if topDiscard.IsNumber() {
		numIx := int(topDiscard)
		if numIx < len(player.cards) && player.cards[numIx].IsHidden() {
			game.TakeDiscard()
			player.Play(topDiscard, game)
			return
		}
	}

	player.Play(game.TakeDeck(), game)
}

// Play applies a single card to the player's hand
func (player *Player) Play(card Card, game *Game) {
	println("playing: " + card.String())
	if card.IsNumber() {
		if int(card) < len(player.cards) && (player.cards[card].IsHidden() || player.cards[card] == CardWild) {
			newCard := player.cards[card].Show()
			player.cards[card] = card
			player.Play(newCard, game)
			return
		}
	} else if card == CardWild {
		for i := 0; i < len(player.cards); i++ {
			if player.cards[i].IsHidden() {
				newCard := player.cards[i].Show()
				player.cards[i] = card
				player.Play(newCard, game)
				return
			}
		}
	} else if card == CardRacoon {
		for ix, dcard := range game.discards {
			if int(dcard) < len(player.cards) && player.cards[dcard].IsHidden() {
				game.discards = append(game.discards[:ix], game.discards[ix:]...)
				game.Discard(card)
				player.Play(dcard, game)
				return
			}
		}
	} else if card == CardElephant || card == CardGiraffe {
		// do nothing
	} else if card == CardSkip {
		var targetIx = -1
		targetPower := -1
		for oIx, opponent := range game.Players {
			if opponent.id == player.id {
				continue
			}
			power := playerPower(opponent)
			if power > targetPower || (power == targetPower && game.rnd.Intn(2) < 1) {
				targetIx = oIx
				targetPower = power
			}
		}
		if targetIx >= 0 {
			printf("Skipping: player %d\n", targetIx)
			game.Players[targetIx].hasSkip = true
			return // the target will discard the skip on their turn
		}
	} else if card == CardDoNotDisturb {
		var target *Player
		targetPower := -1
		targetShowingIx := -1
		for _, opponent := range game.Players {
			if opponent.id == player.id {
				continue
			}
			numShowing := 0
			lastShowingIx := -1
			for oCardIx, oCard := range opponent.cards {
				if !oCard.IsHidden() {
					numShowing++
					lastShowingIx = oCardIx
				}
			}
			if numShowing == 0 {
				continue
			}
			power := playerPower(opponent)

			if power > targetPower || (power == targetPower && game.rnd.Intn(2) < 1) {
				target = opponent
				targetPower = power
				targetShowingIx = lastShowingIx
			}
		}
		if target != nil {
			printf("DnD: player %d, card %d\n", target.id, targetShowingIx)
			target.cards[targetShowingIx] |= CardHidden
		}
	} else {
		println("I don't know how to play this yet: " + card.String())
	}
	game.Discard(card)
}
