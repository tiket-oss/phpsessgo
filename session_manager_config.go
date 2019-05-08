package phpsessgo

import (
	"time"
)

type SessionManagerConfig struct {
	Expiration     time.Duration
	CookiePath     string
	CookieHttpOnly bool
	CookieDomain   string
}
