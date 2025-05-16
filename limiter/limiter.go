package limiter

type Limiter interface {
	Allow(userID string) bool
}
