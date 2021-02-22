package tools

import (
	"math/rand"
	"time"
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func RandomInt(n int) string {
	timeUnix := time.Now().Unix()
	timestamp := time.Unix(timeUnix, 0).Format("20060102150405")
	var letters = []rune(timestamp)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, l)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}
