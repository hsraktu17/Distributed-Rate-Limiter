package limiter

import (
	"sync"
	"time"
)

type tokenBucket struct {
	rate      int
	burst     int
	tokens    int
	lastCheck time.Time
	mu        sync.Mutex
}

func (tb *tokenBucket) allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastCheck).Seconds()
	tb.lastCheck = now

	tb.tokens += int(elapsed * float64(tb.rate))
	if tb.tokens > tb.burst {
		tb.tokens = tb.burst
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

type MemoryLimiter struct {
	rate    int
	burst   int
	buckets map[string]*tokenBucket
	mu      sync.Mutex
}

func NewMemoryLimiter(rate, burst int) *MemoryLimiter {
	return &MemoryLimiter{
		rate:    rate,
		burst:   burst,
		buckets: make(map[string]*tokenBucket),
	}
}

func (ml *MemoryLimiter) Allow(userID string) bool {
	ml.mu.Lock()
	tb, exists := ml.buckets[userID]
	if !exists {
		tb = &tokenBucket{
			rate:      ml.rate,
			burst:     ml.burst,
			tokens:    ml.burst,
			lastCheck: time.Now(),
		}
		ml.buckets[userID] = tb
	}
	ml.mu.Unlock()

	return tb.allow()
}
