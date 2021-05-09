package uniq

import (
	"math/rand"
	"time"
)

var (
	seedValue = time.Now().UnixNano()
	seed      = rand.New(rand.NewSource(seedValue))
)

func GetSeedValue() int64 {
	return seedValue
}

func Int63() int64 {
	return seed.Int63()
}

func Int63n(n int64) int64 {

	return seed.Int63n(n)
}
func Int31() int32 {

	return seed.Int31()
}

func Int31n(n int32) int32 {

	return seed.Int31n(n)
}

func Int() int {

	return seed.Int()
}
