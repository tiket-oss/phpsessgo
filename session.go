package phpsessgo

import "github.com/yvasiyarov/php_session_decoder"

// Session handle creation/modification of session parametr
type Session struct {
	sid   string
	Value php_session_decoder.PhpSession
}

// NewSession create new instance of Session
func NewSession(sid string) *Session {
	return &Session{
		sid: sid,
	}
}

// SessionID return current session ID
func (s *Session) SessionID() string {
	return s.sid
}

// Save session to session source
func (s *Session) Save() error {
	return nil
}
