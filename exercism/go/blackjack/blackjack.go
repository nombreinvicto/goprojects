package blackjack

// declare the card value mapping
var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn,
// given two cards of the player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	totalValue := ParseCard(card1) + ParseCard(card2)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case totalValue <= 11:
		return "H"
	case totalValue >= 12 && totalValue <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	case totalValue >= 17 && totalValue <= 20:
		return "S"
	case totalValue == 21:
		if dealerCard != "ace" && ParseCard(dealerCard) != 10 {
			return "W"
		}
		return "S"
	default:
		return ""
	}
}
