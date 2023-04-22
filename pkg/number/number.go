package number

import (
	"math/rand"
	"time"
)

// GetRandomInt: Return a random number within the min and max range.
// [min, max]
func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GetRandomFloat: Return a random Float64 within the min and max range.
// [min, max]
func GetRandomFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
