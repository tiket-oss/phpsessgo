package phpsessgo

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisSessionHandler session management using redis
type RedisSessionHandler struct {
	SessionHandler
	Client         *redis.Client
	RedisKeyPrefix string
}

// Close the resource
func (h *RedisSessionHandler) Close() {
	if h.Client != nil {
		h.Client.Close()
	}
}

func (h *RedisSessionHandler) Read(sessionID string) (data string, err error) {
	data, err = h.Client.Get(h.sessionRedisKey(sessionID)).Result()
	if err != nil && err.Error() == "redis: nil" {
		data = ""
		err = nil
	}
	return
}

func (h *RedisSessionHandler) Write(sessionID string, sessionData string) error {
	err := h.Client.Set(h.sessionRedisKey(sessionID), sessionData, 0).Err()
	return err
}

func (h *RedisSessionHandler) sessionRedisKey(sessionID string) string {
	return fmt.Sprintf("%s%s", h.RedisKeyPrefix, sessionID)
}
