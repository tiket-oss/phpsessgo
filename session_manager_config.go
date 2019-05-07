package phpsessgo

import (
	"time"
)

type SessionManagerConfig struct {
	Expiration     time.Duration
	CookieHttpOnly bool
	CookieDomain   string
}
