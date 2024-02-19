package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat32 generates a random real between min and max
func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min+1)
}

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

// RandomOwner generates a random owner name
func RandomName(n int) string {
	return RandomString(n)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomInteger generates a random interger number
func RandomInteger(min, max int64) int64 {
	return RandomInt(0, 1000)
}

// RandomLatitude generates a random latitude
func RandomLatitude(min, max float32) float32 {
	return RandomFloat32(min, max)
}

// RandomLongitude generates a random longitude
func RandomLongitude(min, max float32) float32 {
	return RandomFloat32(min, max)
}

// RandomReal generates a random real type number
func RandomReal(min, max float32) float32 {
	return RandomFloat32(min, max)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"PKR", "SAR", "AED", "OMR", "USD", "EUR", "GBP", "CAD",  "CNY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

