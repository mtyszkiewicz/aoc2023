package day06

import (
	"math"

	"github.com/mihonen/polynomials"
)

type Scoreboard struct {
	Times     []int `parser:"'Time' ':' (@Int+)*"`
	Distances []int `parser:"'Distance' ':' (@Int+)*"`
}

func CountWinOptions(totalTime int, bestDistance int) int {
	// In order to match the best times you need to hold the button for x time,
	// for the calculation of which we can use the quadratic equation x^2 - t*x + d = 0,
	// where t is the available time and d is target distance to beat.
	poly := polynomials.CreatePolynomial(1, float64(-totalTime), float64(bestDistance))
	poly.SolveMode = polynomials.DurandKerner
	roots, _ := poly.RealRoots()

	minTime := math.Ceil(roots[1] + 0.001)
	maxTime := math.Floor(roots[0] - 0.001)
	// Offsets +/- 0.001 makes the distance slightly better than a draw.
	return int(maxTime - minTime + 1)
}
