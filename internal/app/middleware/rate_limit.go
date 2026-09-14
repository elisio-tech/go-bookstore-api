package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ipBucket struct {
	count      int
	windowStart time.Time
}

type ipRateLimiter struct {
	mu         sync.Mutex
	requests   int
	window     time.Duration
	buckets    map[string]*ipBucket
	lastClean  time.Time
}

func NewIPRateLimiter(requestsPerWindow int, window time.Duration) *ipRateLimiter {
	l := &ipRateLimiter{
		requests:  requestsPerWindow,
		window:    window,
		buckets:   make(map[string]*ipBucket),
		lastClean: time.Now(),
	}
	go l.startCleanup()
	return l
}

func (l *ipRateLimiter) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	b, ok := l.buckets[ip]
	if !ok || now.Sub(b.windowStart) >= l.window {
		l.buckets[ip] = &ipBucket{count: 1, windowStart: now}
		return true
	}

	if b.count >= l.requests {
		return false
	}
	b.count++
	return true
}

func (l *ipRateLimiter) startCleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		l.mu.Lock()
		for ip, b := range l.buckets {
			if time.Since(b.windowStart) > l.window {
				delete(l.buckets, ip)
			}
		}
		l.lastClean = time.Now()
		l.mu.Unlock()
	}
}

func (l *ipRateLimiter) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if !l.allow(ip) {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "muitas requisições, tente novamente mais tarde"})
			return
		}
		ctx.Next()
	}
}