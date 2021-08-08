package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "0123456789"

var length uint = 24

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randStringFromCharset(length uint, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandString returns random string from defined charset
func RandString() string {
	return randStringFromCharset(length, charset)
}
