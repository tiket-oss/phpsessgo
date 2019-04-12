package phpsessgo

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	RedisSessionName = "PHPREDIS_SESSION"
)

// RedisSessionHandler session management using redis
type RedisSessionHandler struct {
	SessionHandler
	Client      *redis.Client
	SessionName string
}

// Close the resource
func (h *RedisSessionHandler) Close() {
	if h.Client != nil {
		h.Client.Close()
	}
}

func (h *RedisSessionHandler) Read(sessionID string) (data string, err error) {
	data, err = h.Client.Get(sessionRedisKey(sessionID)).Result()
	if err != nil && err.Error() == "redis: nil" {
		data = ""
		err = nil
	}
	return
}

func (h *RedisSessionHandler) Write(sessionID string, sessionData string) error {
	err := h.Client.Set(sessionRedisKey(sessionID), sessionData, 0).Err()
	return err
}

func sessionRedisKey(sessionID string) string {
	return fmt.Sprintf("%s:%s", RedisSessionName, sessionID)
}
