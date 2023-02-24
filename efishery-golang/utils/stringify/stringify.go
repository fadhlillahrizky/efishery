package stringify

import (
	"math/rand"
	"time"
)

func GenRandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	charLength := len(chars)
	rand.Seed(time.Now().UnixNano())
	result := ""
	for i := 0; i < length; i++ {
		result += string(chars[rand.Intn(charLength)])
	}
	return result
}
