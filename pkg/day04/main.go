package day04

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type ScratchCard struct {
	Id      int   `parser:"'Card' @Int':'"`
	Winning []int `parser:"( @Int)* "`
	Drawn   []int `parser:"'|' ( @Int)*"`
}

func (card *ScratchCard) MatchCount() int {
	return len(mapset.NewSet[int](card.Winning...).Intersect(mapset.NewSet[int](card.Drawn...)).ToSlice())
}
