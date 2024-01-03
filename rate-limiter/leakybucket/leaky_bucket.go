package leakybucket

import (
	"context"
	"rhzx3519.github.io/systemdesign/ratelimiter/redis"
)

// RateLimitUsingLeakyBucket .
func RateLimitUsingLeakyBucket(
	userID string,
	uniqueRequestID string,
	intervalInSeconds int64,
	maximumRequests int64,
	ctx context.Context) bool {

	// userID can be apikey, location, ip
	requestCount := redis.RedisClient.LLen(ctx, userID).Val()

	if requestCount >= maximumRequests {
		// drop request
		return false
	}

	// add request id to the end of request queue
	redis.RedisClient.RPush(ctx, userID, uniqueRequestID)

	return true
}
