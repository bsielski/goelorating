package goelorating

import (
	"fmt"
	"math/rand"
	"math/big"
	"strconv"
)

// Interface

type RatingDeltaCalculating interface {
    Calculate(float64, float64) int
}

// Class

type RatingDelta struct {
	kFactor int
	expectedScore ExpextedScoreCalculating
}

func NewRatingDelta(kFactor int, expectedScore ExpextedScoreCalculating) *RatingDelta {
	return &RatingDelta{kFactor, expectedScore}
}

func (e *RatingDelta) Calculate(playerARating, playerBRating int, actualScore float64) int {
	bigE := floatToBig(e.expectedScore.Calculate(playerARating, playerBRating))
	bigA := floatToBig(actualScore)
	bigK := intToBig(e.kFactor)	
	return bigToInt(bigK.Mul(bigK, bigE.Sub(bigA, bigE)))
}

func floatToBig(f float64) *big.Float {
	result, _, _ := big.ParseFloat(fmt.Sprintf("%f", f), 10, 200, big.ToZero)
	return result
}

func intToBig(n int) *big.Float {
	result, _, _ := big.ParseFloat(fmt.Sprintf("%d", n), 10, 200, big.ToZero)
	return result
}

func bigToInt(bf *big.Float) int {
	bI, _ := big.NewFloat(0.0).Copy(bf).Int(nil)
	result, _ := strconv.Atoi(bI.String())
	return result
}

// Mock

type RatingDeltaMock struct {
	playerARating, playerBRating int
	actualScore float64
	calculatingResult int
}

func NewRatingDeltaMock(playerARating, playerBRating int, actualScore float64, fakeResult int) *RatingDeltaMock {
	return &RatingDeltaMock{playerARating, playerBRating, actualScore, fakeResult}
}

func (e *RatingDeltaMock) Calculate(playerARating, playerBRating int, actualScore float64) int {
	if playerARating == e.playerARating && playerBRating == e.playerBRating && actualScore == e.actualScore {
		return e.calculatingResult
	}
	var unwantedResult int
	for { // because is a chance that random == wanted result
		unwantedResult = rand.Intn(1000) - 500
		if unwantedResult != e.calculatingResult {
			break
		}
	}
	return unwantedResult
}
