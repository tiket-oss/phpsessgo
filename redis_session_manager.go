package phpsessgo

import "github.com/go-redis/redis"

// NewRedisSessionManager create new instance of SessionManager
func NewRedisSessionManager(client *redis.Client, config SessionManagerConfig) SessionManager {
	sessionManager := &sessionManager{
		sessionName: DefaultSessionName,
		sidCreator:  &UUIDCreator{},
		handler: &RedisSessionHandler{
			Client:         client,
			RedisKeyPrefix: DefaultRedisKeyPrefix,
			Expiration:     config.Expiration,
		},
		encoder: &PHPSessionEncoder{},
		config:  config,
	}
	return sessionManager
}
