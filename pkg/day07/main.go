package day07

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

var cardValues = map[rune]string{
	'A': "14",
	'K': "13",
	'Q': "12",
	'J': "11",
	'T': "10",
	'9': "09",
	'8': "08",
	'7': "07",
	'6': "06",
	'5': "05",
	'4': "04",
	'3': "03",
	'2': "02",
}

var cardValuesJockersAllowed = map[rune]string{
	'A': "14",
	'K': "13",
	'Q': "12",
	'T': "10",
	'9': "09",
	'8': "08",
	'7': "07",
	'6': "06",
	'5': "05",
	'4': "04",
	'3': "03",
	'2': "02",
	'J': "01",
}

const (
	five_of_kind  = 6
	four_of_kind  = 5
	full_house    = 4
	three_of_kind = 3
	two_pair      = 2
	one_pair      = 1
	high_card     = 0
)

type Player struct {
	Hand     string
	Bid      int
	ScoreOne int
	ScoreTwo int
}

func NewPlayer(hand string, bid int, jockersAllowed bool) *Player {
	return &Player{
		Hand:     hand,
		Bid:      bid,
		ScoreOne: ScoreHandRuleOne(hand, jockersAllowed),
		ScoreTwo: ScoreHandRuleTwo(hand, jockersAllowed),
	}
}

func ComparePlayers(p1 *Player, p2 *Player) int {
	result := cmp.Compare(p1.ScoreOne, p2.ScoreOne)
	if result != 0 {
		return result
	}
	return cmp.Compare(p1.ScoreTwo, p2.ScoreTwo)
}

func ScoreHandRuleOne(hand string, jockersAllowed bool) int {
	frequency := make(map[rune]int)
	for _, char := range hand {
		frequency[char] = frequency[char] + 1
	}
	counts := []int{}
	for _, value := range frequency {
		counts = append(counts, value)
	}
	jockerCount := frequency['J']
	jockerPresent := jockerCount != 0

	switch {
	case slices.Contains(counts, 5): // five of kind
		return five_of_kind
	case slices.Contains(counts, 4): // four of kind
		if jockersAllowed && jockerPresent {
			return five_of_kind
		}
		return four_of_kind
	case slices.Contains(counts, 3): // full house or 3 of a kind
		if len(counts) == 2 {
			if jockersAllowed && jockerPresent {
				return five_of_kind
			}
			return full_house
		} else {
			if jockersAllowed && jockerPresent {
				return four_of_kind
			}
			return three_of_kind
		}
	case slices.Contains(counts, 2): // two pair or one pair
		if len(counts) == 3 {
			if jockersAllowed {
				if jockerCount == 2 {
					return four_of_kind
				}
				if jockerCount == 1 {
					return full_house
				}
			}
			return two_pair
		} else {
			if jockersAllowed && jockerPresent {
				return three_of_kind
			}
			return one_pair
		}
	default:
		if jockersAllowed && jockerPresent {
			return one_pair
		}
		return high_card
	}
}

func ScoreHandRuleTwo(hand string, jockersAlloweds bool) int {
	valueMap := cardValues
	if jockersAlloweds {
		valueMap = cardValuesJockersAllowed
	}
	v1 := valueMap[rune(hand[0])]
	v2 := valueMap[rune(hand[1])]
	v3 := valueMap[rune(hand[2])]
	v4 := valueMap[rune(hand[3])]
	v5 := valueMap[rune(hand[4])]
	result, _ := strconv.Atoi(fmt.Sprintf("%s%s%s%s%s", v1, v2, v3, v4, v5))
	return result
}
