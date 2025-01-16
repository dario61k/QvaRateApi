package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type requestInfo struct {
	count    int
	lastSeen time.Time
}

var requestCount = make(map[string]*requestInfo)
var mutex sync.Mutex

func RateLimit(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {

		clientIp := c.ClientIP()

		mutex.Lock()

		info, exist := requestCount[clientIp]
		if !exist {
			requestCount[clientIp] = &requestInfo{
				count:    1,
				lastSeen: time.Now(),
			}
		} else {
			if time.Since(info.lastSeen) > duration {
				info.count = 1
				info.lastSeen = time.Now()
			} else {
				info.count++
			}
		}

		if requestCount[clientIp].count > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			c.Abort()
			mutex.Unlock()
			return
		}

		mutex.Unlock()

		c.Next()
	}
}
