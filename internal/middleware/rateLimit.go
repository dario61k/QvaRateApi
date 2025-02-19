package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type requestInfo struct {
	count       int
	lastSeen    time.Time
	isBlocked   bool
	blockExpiry time.Time
}

var requestCount = make(map[string]*requestInfo)
var mutex sync.Mutex

func RateLimit(limit int, duration, blockDuration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		mutex.Lock()
		info, exists := requestCount[clientIP]

		if !exists {
			requestCount[clientIP] = &requestInfo{
				count:    1,
				lastSeen: time.Now(),
			}
		} else {

			if info.isBlocked {
				if time.Now().After(info.blockExpiry) {

					info.isBlocked = false
					info.count = 1
					info.lastSeen = time.Now()
				} else {

					c.JSON(http.StatusTooManyRequests, gin.H{
						"error": "Too Many Requests - You are temporarily blocked",
					})
					mutex.Unlock()
					c.Abort()
					return
				}
			} else {

				if time.Since(info.lastSeen) > duration {
					info.count = 1
					info.lastSeen = time.Now()
				} else {
					info.count++
				}

				if info.count > limit {
					info.isBlocked = true
					info.blockExpiry = time.Now().Add(blockDuration)
					c.JSON(http.StatusTooManyRequests, gin.H{
						"error": "Too Many Requests - You are temporarily blocked",
					})
					mutex.Unlock()
					c.Abort()
					return
				}
			}
		}

		mutex.Unlock()
		c.Next()
	}
}
