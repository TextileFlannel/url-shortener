package storage

import (
	"math/rand"
	"sync"
	"time"
)

const (
	shortURLLength = 10
	charset        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

var (
	seededRand *rand.Rand
	once       sync.Once
)

func initGenerator() {
	once.Do(func() {
		seededRand = rand.New(
			rand.NewSource(time.Now().UnixNano()))
	})
}

func GenerateShortURL() string {
	initGenerator()

	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
