package deck

import (
	"fmt"
)

//Suit definition

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

//Rank definition

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
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
)

var ranks = [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func (r Rank) String() string {
	return ranks[r-1] //[...] - specifies the length is equal to the number of elements in the array literal
}

//Card definition

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank, c.Suit)
}
