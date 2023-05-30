package cache

import (
	"math/rand"
	"time"
)

const devitation = 0.05

func unstableDuration(base time.Duration) time.Duration {
	return time.Duration(float64(base) * (1 + devitation*(rand.Float64()-0.5)))
}

func unstableInt(base int64) int64 {
	return int64(float64(base) * (1 + devitation*(rand.Float64()-0.5)))
}
