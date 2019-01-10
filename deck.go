package deck

import (
	"math/rand"
	"sort"
	"time"
)

//Deck definition and creation

type Deck []Card

//...func([]Card) []Card -> 0 or more parameters of type function that is able to take in a a slice of cards(deck of cards) and return a slice of cards, used for adding new cards to the deck
func New(opts ...func(Deck) Deck) Deck {
	var cards Deck

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(ranks); j++ {
			cards = append(cards, Card{Suit(i), Rank(j + 1)})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func (deck *Deck) DealCard() Card {
	var card Card
	card, *deck = (*deck)[0], (*deck)[1:]
	return card
}

//Deck options

func Sort(less func(cards Deck) func(i, j int) bool) func(Deck) Deck {
	return func(cards Deck) Deck {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards Deck) Deck {
	sort.Slice(cards, Less(cards))
	return cards
}

func ReverseSort(cards Deck) Deck {
	sort.Slice(cards, Reverse(cards))
	return cards
}

func Less(cards Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func Reverse(cards Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(3) + int(c.Rank)
}

func Shuffle(cards Deck) Deck {
	shuffledCards := make(Deck, len(cards))
	//rand.Perm(4) = [3,1,2,0] - a slice of integers from 0 to 3 randomly arranged
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm { //i - index of the permutation, j - value inside the permutation
		shuffledCards[i] = cards[j]
	}

	return shuffledCards
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func ShuffleTestable(cards Deck) Deck {
	shuffledCards := make(Deck, len(cards))
	//rand.Perm(4) = [3,1,2,0] - a slice of integers from 0 to 3 randomly arranged
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm { //i - index of the permutation, j - value inside the permutation
		shuffledCards[i] = cards[j]
	}

	return shuffledCards
}

func Filter(filterFunction func(card Card) bool) func(Deck) Deck {
	return func(cards Deck) Deck {
		var filteredDeck Deck
		for _, card := range cards {
			if !filterFunction(card) { //if the given card doesn't pass the filter function add it to the filtered deck
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}

func Amount(n int) func(Deck) Deck {
	return func(cards Deck) Deck {
		var allDecks Deck
		for i := 0; i < n; i++ {
			allDecks = append(allDecks, cards...) // cards... - unpacks the array to it's elements
		}
		return allDecks
	}
}
