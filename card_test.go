package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Diamonds})
	fmt.Println(Card{Rank: Two, Suit: Clubs})
	fmt.Println(Card{Rank: Five, Suit: Spades})
	fmt.Println(Card{Rank: Ten, Suit: Hearts})
	fmt.Println(Card{Rank: Jack, Suit: Clubs})
	fmt.Println(Card{Rank: King, Suit: Diamonds})

	// Output:
	// AD
	// 2C
	// 5S
	// 10H
	// JC
	// KD
}

func TestNew(t *testing.T) {
	cards := New()
	//13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in deck")
	}

	for suit := 0; suit < len(suits); suit++ {
		for rank := 0; rank < len(ranks); rank++ {
			currentCard := (suit * 13) + rank //0 -> 52
			if cards[currentCard].Suit != Suit(suit) && cards[currentCard].Rank != Rank(rank) {
				t.Error("Wrong cards in deck")
			}
		}
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Clubs, Ace}

	if cards[0] != exp {
		t.Error("Expected AC as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Clubs, Ace}

	if cards[0] != exp {
		t.Error("Expected AC as first card. Received: ", cards)
	}
}

func TestSortReverse(t *testing.T) {
	cards := New(Sort(Reverse))
	exp := Card{Spades, King}

	if cards[0] != exp {
		t.Error("Expected KS as first card. Received: ", cards)
	}
}

func TestReverseSort(t *testing.T) {
	cards := New(ReverseSort)
	exp := Card{Spades, King}
	fmt.Println(int(King))
	fmt.Println(int(cards[0].Rank))
	fmt.Println(cards[0])
	fmt.Println(len(ranks))
	fmt.Println(ranks)
	fmt.Println(exp, cards[0])
	fmt.Println(exp, cards[0])
	fmt.Println(exp == cards[0])
	fmt.Println(exp == cards[0])
	if cards[0] != exp {
		t.Error("Expected KS as first card. Received: ", cards)
	}
}

func TestShuffleTestable(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	cards := New()
	first := cards[40]
	second := cards[35]

	shuffledCards := New(ShuffleTestable)

	if shuffledCards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, shuffledCards[0])
	}
	if shuffledCards[1] != second {
		t.Errorf("Expected the second card to be %s, received %s", second, shuffledCards[1])
	}
}

func TestFilter(t *testing.T) {
	filterFunction := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three //filter out twos and threes
	}

	cards := New(Filter(filterFunction))

	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Amount(3))

	if (len(cards)) != 52*3 {
		t.Errorf("Expected %d cards received %d cards", 52*3, len(cards))
	}
}
