package phpsessgo

import "github.com/go-redis/redis"

// NewRedisSessionManager create new instance of SessionManager
func NewRedisSessionManager(client *redis.Client) *SessionManager {
	sessionManager := &SessionManager{
		SessionName: DefaultSessionName,
		SIDCreator:  &UUIDCreator{},
		Handler: &RedisSessionHandler{
			Client:         client,
			RedisKeyPrefix: DefaultRedisKeyPrefix,
		},
		Encoder: &PHPSessionEncoder{},
	}
	return sessionManager
}
