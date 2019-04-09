package phpsessgo

import "net/http"

// SessionManager handle session creation/modification
type SessionManager struct {
	SessionName string
	SIDCreator  SessionIDCreator
	Handler     SessionHandler
}

// NewSessionManager create new instance of SessionManager
func NewSessionManager(config SessionConfig) (*SessionManager, error) {
	sidCreator := NewSessionIDCreator()
	sessionManager := &SessionManager{
		SessionName: config.Name,
		SIDCreator:  sidCreator,
	}
	return sessionManager, nil
}

// Start is adoption of PHP start_session() to return current active session
func (m *SessionManager) Start(w http.ResponseWriter, r *http.Request) Session {
	sid := m.getSIDFromHTTPCookie(r.Cookies())
	if sid == "" {
		sid = m.SIDCreator.CreateSID()
		m.setSIDToHTTPCookie(w, sid)
	}
	return NewSession(sid)
}

func (m *SessionManager) getSIDFromHTTPCookie(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == m.SessionName {
			return cookie.Value
		}
	}
	return ""
}

func (m *SessionManager) setSIDToHTTPCookie(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{
		Name:  m.SessionName,
		Value: sid,
	})
}
