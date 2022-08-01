package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/jitter"
	"github.com/Rican7/retry/strategy"
)

func main() {
	t0 := time.Now()

	action := func(attempt uint) error {
		diff := time.Now().Sub(t0)
		fmt.Println("attempt:", attempt, " internal diff:", diff)
		return errors.New("something happened")
	}

	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	retry.Retry(
		action,
		strategy.Limit(16),
		strategy.BackoffWithJitter(
			backoff.Fibonacci(80*time.Millisecond),
			jitter.Deviation(random, 0.5),
		),
	)
	diff := time.Now().Sub(t0)
	fmt.Println("total diff:", diff)
}
