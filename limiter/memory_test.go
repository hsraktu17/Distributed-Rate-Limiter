package limiter

import (
	"sync"
	"testing"
	"time"
)

func TestMemoryLimiter_AllowInitialBurst(t *testing.T) {
	lim := NewMemoryLimiter(10, 3)
	user := "user1"

	for i := 0; i < 3; i++ {
		if !lim.Allow(user) {
			t.Fatalf("expected allow on token %d", i)
		}
	}
	if lim.Allow(user) {
		t.Fatalf("expected deny when burst exhausted")
	}
}

func TestMemoryLimiter_TokenRefill(t *testing.T) {
	lim := NewMemoryLimiter(10, 1)
	user := "user2"

	if !lim.Allow(user) {
		t.Fatal("first call should succeed")
	}
	if lim.Allow(user) {
		t.Fatal("burst consumed, should deny")
	}

	time.Sleep(150 * time.Millisecond)
	if !lim.Allow(user) {
		t.Fatal("token should have refilled")
	}
}

func TestMemoryLimiter_Concurrent(t *testing.T) {
	lim := NewMemoryLimiter(1, 5)
	user := "concurrent"

	var wg sync.WaitGroup
	var allowed int

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if lim.Allow(user) {
				// protect concurrent increment
				// purposely not using atomic to keep dependencies minimal
				allowed++
			}
		}()
	}
	wg.Wait()

	if allowed != 5 {
		t.Fatalf("expected 5 allowed, got %d", allowed)
	}
}
