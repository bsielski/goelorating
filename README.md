# Calculate Elo rating in Go (Go language, not Go game)

It has just two objects:

* `ExpectedScore` - with one public method `Calculate(ratingA, ratingB int) expectedScoreForA float64`.
* `RatingDelta` - with one public method `Calculate(ratingA, ratingB int, actualScore float64) ratingDeltaForA int`

## Sample usage

```
package main

import (
	"fmt"
	elo "github.com/bsielski/goelorating"
)

func main() {
	kFactor := 32
	deviation := 400
	expectedScore := elo.NewExpectedScore(deviation)
	playerARating := 1678
	playerBRating := 1432
	actualScore := 0.6

	// If you want expected score for player A

	fmt.Println(expectedScore.Calculate(playerARating, playerBRating))

	// If you want rating delta for player A

	ratingDelta := elo.NewRatingDelta(kFactor, expectedScore)
	fmt.Println(ratingDelta.Calculate(playerARating, playerBRating, actualScore)) 
}
```