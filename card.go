package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

var suits = [...]string{"C", "D", "H", "S"}

func (s Suit) String() string {
	return suits[s] //[...] - specifies the length is equal to the number of elements in the array literal
}

type Rank uint8

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var ranks = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func (r Rank) String() string {
	return ranks[r] //[...] - specifies the length is equal to the number of elements in the array literal
}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}

//...func([]Card) []Card -> 0 or more parameters of type function that is able to take in a a slice of cards(deck of cards) and return a slice of cards, used for adding new cards to the deck
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(ranks); j++ {
			cards = append(cards, Card{Suit(i), Rank(j)})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func ReverseSort(cards []Card) []Card {
	sort.Slice(cards, Reverse(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func Reverse(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(3) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	shuffledCards := make([]Card, len(cards))
	//rand.Perm(4) = [3,1,2,0] - a slice of integers from 0 to 3 randomly arranged
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm { //i - index of the permutation, j - value inside the permutation
		shuffledCards[i] = cards[j]
	}

	return shuffledCards
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

func ShuffleTestable(cards []Card) []Card {
	shuffledCards := make([]Card, len(cards))
	//rand.Perm(4) = [3,1,2,0] - a slice of integers from 0 to 3 randomly arranged
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm { //i - index of the permutation, j - value inside the permutation
		shuffledCards[i] = cards[j]
	}

	return shuffledCards
}

func Filter(filterFunction func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var filteredDeck []Card
		for _, card := range cards {
			if !filterFunction(card) { //if the given card doesn't pass the filter function add it to the filtered deck
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var allDecks []Card
		for i := 0; i < n; i++ {
			allDecks = append(allDecks, cards...) // cards... - unpacks the array to it's elements
		}
		return allDecks
	}
}
