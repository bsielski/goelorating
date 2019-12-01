package goelorating

import (
	"math/rand"
	"math"
)

// Interface

type ExpextedScoreCalculating interface {
    Calculate(int, int) float64
}

// Class

type ExpectedScore struct {
	deviation int
}

func NewExpectedScore(deviation int) *ExpectedScore {
	return &ExpectedScore{deviation}
}

func (e *ExpectedScore) Calculate(playerARating, playerBRating int) float64 {
	return 1 / (1 + math.Pow(10, float64(playerBRating - playerARating) / float64(e.deviation)))
	//return rand.Float64()
}

// Mock

type ExpectedScoreMock struct {
	playerARating, playerBRating int
	calculatingResult float64
}

func NewExpectedScoreMock(playerARating, playerBRating int, fakeResult float64) *ExpectedScoreMock {
	return &ExpectedScoreMock{playerARating, playerBRating, fakeResult}
}

func (e *ExpectedScoreMock) Calculate(playerARating, playerBRating int) float64 {
	if playerARating == e.playerARating && playerBRating == e.playerBRating {
		return e.calculatingResult
	}
	var unwantedResult float64
	for { // because "is" a chance that random == wanted result
		unwantedResult = rand.Float64()
		if unwantedResult != e.calculatingResult {
			break
		}
	}
	return unwantedResult

}
