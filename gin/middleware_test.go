package ginratelimiter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/limiter"
)

func setupRouter(l limiter.Limiter) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Middleware(l))
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	return r
}

func TestMiddleware_MissingUserID(t *testing.T) {
	r := setupRouter(limiter.NewMemoryLimiter(1, 1))

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestMiddleware_RateLimit(t *testing.T) {
	lim := limiter.NewMemoryLimiter(10, 2)
	r := setupRouter(lim)

	req := httptest.NewRequest("GET", "/ping?user_id=u1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 first request, got %d", w.Code)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 second request, got %d", w.Code)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429 after limit, got %d", w.Code)
	}
}
