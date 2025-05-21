package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "0123456789"
)

func GenerateRandomVehicleId() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	prefixLen := random.Intn(2) + 1
	numberLen := random.Intn(4) + 1
	suffixLen := random.Intn(3) + 1

	prefix := generateRandomString(random, prefixLen, letters)
	number := generateRandomString(random, numberLen, digits)
	suffix := generateRandomString(random, suffixLen, letters)

	return prefix + number + suffix
}

func generateRandomString(r *rand.Rand, length int, providedString string) string {
	b := strings.Builder{}

	b.Grow(length)

	for i := 0; i < length; i++ {
		b.WriteByte(providedString[r.Intn(len(providedString))])
	}

	return b.String()
}
