package phpsessgo

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

// SessionManager handle session creation/modification
type SessionManager struct {
	SessionName string
	SIDCreator  SessionIDCreator
	Handler     SessionHandler
}

// NewSessionManager create new instance of SessionManager
func NewSessionManager(config SessionConfig) (*SessionManager, error) {

	sessionManager := &SessionManager{
		SessionName: config.Name,
		SIDCreator:  &sessionIDCreator{},
		Handler: &RedisSessionHandler{
			Client: redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			}),
			SessionName: config.Name,
		},
	}
	return sessionManager, nil
}

// Start is adoption of PHP start_session() to return current active session
func (m *SessionManager) Start(w http.ResponseWriter, r *http.Request) (session *Session, err error) {
	sid := m.getFromCookies(r.Cookies())
	if sid == "" {
		sid = m.SIDCreator.CreateSID()
		m.setToCookies(w, sid)

		session = NewSession(sid)
		return
	}

	if m.Handler != nil {
		data, _ := m.Handler.Read(sid)
		fmt.Println(data)
	}

	session = NewSession(sid)
	return
}

func (m *SessionManager) getFromCookies(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == m.SessionName {
			return cookie.Value
		}
	}
	return ""
}

func (m *SessionManager) setToCookies(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{
		Name:  m.SessionName,
		Value: sid,
	})
}
