package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomDescription(n int) string {
	return faker.Sentence()
}

// GenerateRandomPrice generates a random price and converts it to pgtype.Numeric
func RandomPrice(min, max float64) (pgtype.Numeric, error) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random price
	price := min + rand.Float64()*(max-min)

	// Convert the price to a string with two decimal places
	priceStr := fmt.Sprintf("%.2f", price)

	// Create a pgtype.Numeric instance and set the value
	var numeric pgtype.Numeric
	if err := numeric.Scan(priceStr); err != nil {
		return numeric, err
	}

	return numeric, nil
}

func RandomStock() int32 {
	return rand.Int31n(1000)
}
