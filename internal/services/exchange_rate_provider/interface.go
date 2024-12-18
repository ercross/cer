package exchange

import (
	"math/rand"
	"time"
)

type RateProvider interface {
	FetchExchangeRate(pair string) (float64, error)
	Name() string
}

// sleep for a random duration, usually for maxSleepMilliseconds or lesser duration
func sleep() {
	maxSleepMilliseconds := 2000
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// random sleep duration between 1 and maxSleepMilliseconds
	sleepDuration := rand.Intn(maxSleepMilliseconds) + 1
	time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
}
