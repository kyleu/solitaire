package deck_test

import (
	"testing"

	"github.com/kyleu/solitaire/app/game/deck"
)

func getFreshDeck(faceUp bool) *deck.Deck {
	return deck.FreshDeck(1, nil, nil, faceUp, 0)
}

func TestEqual(t *testing.T) {
	t.Parallel()

	if !getFreshDeck(false).Equals(getFreshDeck(false)) {
		t.Error("fresh decks are not equal")
	}
	if !getFreshDeck(true).Equals(getFreshDeck(true)) {
		t.Error("fresh face up decks are not equal")
	}
	if getFreshDeck(true).Equals(getFreshDeck(false)) {
		t.Error("fresh decks that differ in faceUp are erroneously equal")
	}

	if !getFreshDeck(true).EqualsSimple(getFreshDeck(true)) {
		t.Error("fresh decks that differ in faceUp are not simply equal")
	}
	if !getFreshDeck(true).EqualsSimple(getFreshDeck(false)) {
		t.Error("fresh decks that differ in faceUp are not simply equal")
	}
}
