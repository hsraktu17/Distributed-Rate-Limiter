package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	rate      int
	burst     int
	token     int
	lastCheck time.Time
	mu        sync.Mutex
}

func NewLimiter(rate, burst int) *Limiter {
	return &Limiter{
		rate:      rate,
		burst:     burst,
		token:     burst,
		lastCheck: time.Now(),
	}
}
