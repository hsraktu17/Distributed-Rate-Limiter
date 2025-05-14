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

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastCheck).Seconds()
	l.lastCheck = now

	l.token += int(elapsed * float64(l.rate))
	if l.token > l.burst {
		l.token = l.burst
	}

	if l.token > 0 {
		l.token--
		return true
	}
	return false
}
