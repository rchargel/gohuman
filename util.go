package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math"
	mtrand "math/rand"
	"time"
)

const idLength = 15

// SHA512 Creates a Base64 encoded Message Digest using SHA-512
func SHA512(message string) string {
	c := sha512.New()
	c.Write([]byte(message))
	hash := c.Sum(nil)
	return base64.URLEncoding.EncodeToString(hash)
}

// NewID creates a random string for use in identifiers.
// xxxxxyyxxxxxmmxxxxxdd
func NewID(data ...string) string {
	now := time.Now()
	random := make([]byte, 32)
	rand.Read(random)
	pre := base64.URLEncoding.EncodeToString(random)

	for _, str := range data {
		pre += str
	}
	hash := SHA512(pre)
	l := len(hash)
	s := mtrand.Intn(l - idLength)
	f := fmt.Sprintf("%02d%v%02d%v%02d%v%02d", int(math.Abs(float64(now.Hour()^now.Minute()^now.Second()))),
		hash[s:s+5], now.Year()-2000, hash[s+5:s+10], now.Month(), hash[s+10:s+15], now.Day())
	return f
}
