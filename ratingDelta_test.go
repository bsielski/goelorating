package goelorating

import (
	"testing"
)

func TestRatingDelta(t *testing.T) {
	cases := []struct {
		kFactor int
		playerARating, playerBRating int
		expectedScore float64
		actualScore float64
		delta int
	}{
		{10, 2000, 1939, 0.452, 0.3,  -1},
		{20, 1410, 1863, 1.0,   0.5, -10},
		{40, 1920, 1917, 0.952, 1.0,   1},
		{32, 1980, 1981, 0.02,  0.7,  21},

	}
	for i, c := range cases {
		expectedScoreMock := NewExpectedScoreMock(c.playerARating, c.playerBRating, c.expectedScore)
		delta := NewRatingDelta(c.kFactor, expectedScoreMock).Calculate(c.playerARating, c.playerBRating, c.actualScore)
		if delta != c.delta {
			t.Errorf(
				"Case %d: output %d, but should be %d",
				i,
				delta,
				c.delta,
			)
		}
	}
}

func TestRatingDeltaMock(t *testing.T) {
	cases := []struct {
		expectedPlayerARating, expectedPlayerBRating int
		expectedActualScore float64
		expectedResult int
		passedPlayerARating, passedPlayerBRating int
		passedActualScore float64
		shouldReturnAsExpected bool
	}{
		{2000, 2000, 1.0,     323,   2000, 2000, 1.0,      true},
		{1200, 1000, 0.32100,  20,   1200, 1000, 0.2100,  false},
		{1700, 1900, 0.612,     0,   1700, 1240, 0.612,   false},
		{1150, 1900, 0.2100, 1000,   1250, 1901, 0.110,   false},
		{1150, 1900, 0.999,  2102,   1900, 1150, 0.999,  false},
	}
	for i, c := range cases {
		mock := NewRatingDeltaMock(c.expectedPlayerARating, c.expectedPlayerBRating, c.expectedActualScore, c.expectedResult)
		result := mock.Calculate(c.passedPlayerARating, c.passedPlayerBRating, c.passedActualScore)
		isReturnAsExpected := c.expectedResult == result
		
		if isReturnAsExpected != c.shouldReturnAsExpected  {
			t.Errorf(
				"Case %d: mock correctness is: %t, but should be %t",
				i,
				isReturnAsExpected,
				c.shouldReturnAsExpected,
			)
		}
	}
}
