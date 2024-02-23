package Dealer

import (
	"errors"
	"math/rand"
)

var deck []string
var dealtCards []string

func AddCardsToDeck(cardsToAdd []string) {
	deck = append(deck, cardsToAdd...)
}

func AddDealtCardsBackToDeck() {
	AddCardsToDeck(dealtCards)
	dealtCards = dealtCards[:0]
}

func removeElementFromSlice(slice []string, index int) ([]string, error) {
	if index >= len(slice) || index < 0 {
		return nil, errors.New("cannot remove element from slice - index out of bounds")
	} else {
		return append(slice[:index], slice[index+1:]...), nil
	}
}

func ShuffleDeck() {
	var i int
	var err error
	currentDeck := deck[:]
	var shuffledDeck []string
	for len(deck) > 0 {
		i = rand.Intn(len(deck) - 1)
		shuffledDeck = append(shuffledDeck, currentDeck[i])
		currentDeck, err = removeElementFromSlice(currentDeck, i)
		if err != nil {
			panic(err)
		}
	}
	deck = shuffledDeck
}

func DealCard() (string, error) {
	if len(deck) == 0 {
		errorMessage := "cannot deal card - empty deck..."
		return errorMessage, errors.New(errorMessage)
	} else {
		card := deck[0]
		var err error
		dealtCards = append(dealtCards, card)
		deck, err = removeElementFromSlice(deck, 0)
		if err != nil {
			panic(err)
		}
		return card, nil
	}
}
