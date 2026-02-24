package payspace

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RateLimitInfo holds rate limit data extracted from response headers.
type RateLimitInfo struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

// rateLimitTracker keeps the most recent rate limit state.
type rateLimitTracker struct {
	mu   sync.RWMutex
	info RateLimitInfo
}

func newRateLimitTracker() *rateLimitTracker {
	return &rateLimitTracker{}
}

func (r *rateLimitTracker) update(info RateLimitInfo) {
	r.mu.Lock()
	r.info = info
	r.mu.Unlock()
}

// Current returns the latest known rate limit state.
func (r *rateLimitTracker) Current() RateLimitInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.info
}

// parseRateLimitHeaders extracts rate limit info from HTTP response headers.
func parseRateLimitHeaders(h http.Header) RateLimitInfo {
	info := RateLimitInfo{}
	if v := h.Get("X-RateLimit-Limit"); v != "" {
		info.Limit, _ = strconv.Atoi(v)
	}
	if v := h.Get("X-RateLimit-Remaining"); v != "" {
		info.Remaining, _ = strconv.Atoi(v)
	}
	if v := h.Get("X-RateLimit-Reset"); v != "" {
		if sec, err := strconv.ParseInt(v, 10, 64); err == nil {
			info.Reset = time.Unix(sec, 0)
		}
	}
	return info
}
