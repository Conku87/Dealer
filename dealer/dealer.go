package dealer

import (
	"errors"
	"math/rand"
)

type Dealer struct {
	deck       []string
	dealtCards []string
}

func NewDealer() *Dealer {
	return &Dealer{deck: []string{}, dealtCards: []string{}}
}

func (dealer *Dealer) AddCardsToDeck(cardsToAdd []string) {
	dealer.deck = append(dealer.deck, cardsToAdd...)
}

func (dealer *Dealer) AddDealtCardsBackToDeck() {
	dealer.AddCardsToDeck(dealer.dealtCards)
	dealer.dealtCards = dealer.dealtCards[:0]
}

func removeElementFromSlice(slice []string, index int) ([]string, error) {
	if index >= len(slice) || index < 0 {
		return nil, errors.New("cannot remove element from slice - index out of bounds")
	} else {
		return append(slice[:index], slice[index+1:]...), nil
	}
}

func (dealer *Dealer) ShuffleDeck() {
	var i int
	var err error
	currentDeck := dealer.deck[:]
	var shuffledDeck []string
	for len(currentDeck) > 0 {
		i = rand.Intn(len(currentDeck))
		shuffledDeck = append(shuffledDeck, currentDeck[i])
		currentDeck, err = removeElementFromSlice(currentDeck, i)
		if err != nil {
			panic(err)
		}
	}
	dealer.deck = shuffledDeck
}

func (dealer *Dealer) DealCard() (string, error) {
	if len(dealer.deck) == 0 {
		errorMessage := "cannot deal card - empty deck..."
		return errorMessage, errors.New(errorMessage)
	} else {
		card := dealer.deck[0]
		var err error
		dealer.dealtCards = append(dealer.dealtCards, card)
		dealer.deck, err = removeElementFromSlice(dealer.deck, 0)
		if err != nil {
			panic(err)
		}
		return card, nil
	}
}
