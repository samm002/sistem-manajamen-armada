package util

import (
	"math/rand"
	"time"
)

func GenerateRandomCoordinate(maxValue int) float64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return -float64(maxValue) + random.Float64()*float64(maxValue)
}
