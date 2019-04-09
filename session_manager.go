package phpsessgo

import "net/http"

// SessionManager handle session creation/modification
type SessionManager interface {
	Retrieve(w http.ResponseWriter, r *http.Request) Session
}

// NewSessionManager create new instance of SessionManager
func NewSessionManager(config SessionConfig) (SessionManager, error) {
	sidGenerator := NewSessionIDGenerator()
	sessionManager := &sessionManager{
		sessionName:  config.Name,
		sidGenerator: sidGenerator,
	}
	return sessionManager, nil
}

type sessionManager struct {
	sessionName  string
	sidGenerator SessionIDGenerator
}

// Retrieve
func (m *sessionManager) Retrieve(w http.ResponseWriter, r *http.Request) Session {
	sid := m.getSIDFromHTTPCookie(r.Cookies())
	if sid == "" {
		sid = m.sidGenerator.CreateSID()
		m.setSIDToHTTPCookie(w, sid)
	}
	return NewSession(sid)
}

func (m *sessionManager) getSIDFromHTTPCookie(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == m.sessionName {
			return cookie.Value
		}
	}
	return ""
}

func (m *sessionManager) setSIDToHTTPCookie(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{
		Name:  m.sessionName,
		Value: sid,
	})
}
