package goelorating

import (
	"fmt"
	"testing"
)

func TestExpectedScore(t *testing.T) {
	cases := []struct {
		deviation int
		playerARating, playerBRating int
		score string
	}{
		{400, 2000, 2000, "0.50"},
		{400, 1200, 1000, "0.76"},
		{400, 1700, 1900, "0.24"},
		{400, 1150, 1900, "0.01"},
	}
	for i, c := range cases {
		unroundedScore := NewExpectedScore(c.deviation).Calculate(c.playerARating, c.playerBRating)
		score := fmt.Sprintf("%.2f", unroundedScore)
		if score != c.score {
			t.Errorf(
				"Case %d: output %s, but should be %s",
				i,
				score,
				c.score,
			)
		}
	}
}

func TestExpectedScoreMock(t *testing.T) {
	cases := []struct {
		expectedARating, expectedBRating int
		expectedResult float64
		passedARating, passedBRating int
		shouldReturnAsExpected bool
	}{
		{3000, 2000, 0.50, 3000, 2000, true},
		{2000, 2000, 0.53, 2000, 1000, false},
		{1200, 1000, 0.06, 1201, 1000, false},
		{1200, 1000, 0.76, 1000, 1200, false},
		{1700, 1900, 0.84, 1002, 2131, false},
	}
	for i, c := range cases {
		mock := NewExpectedScoreMock(c.expectedARating, c.expectedBRating, c.expectedResult)
		result := mock.Calculate(c.passedARating, c.passedBRating)
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
