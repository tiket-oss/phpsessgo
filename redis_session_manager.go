package phpsessgo

import "github.com/go-redis/redis"

// NewRedisSessionManager create new instance of SessionManager
func NewRedisSessionManager(client *redis.Client, config SessionManagerConfig) *SessionManager {
	sessionManager := &SessionManager{
		SessionName: DefaultSessionName,
		SIDCreator:  &UUIDCreator{},
		Handler: &RedisSessionHandler{
			Client:         client,
			RedisKeyPrefix: DefaultRedisKeyPrefix,
			Expiration:     config.Expiration,
		},
		Encoder: &PHPSessionEncoder{},
	}
	return sessionManager
}
