package phpsessgo

import "github.com/imantung/phpsessgo/phpencode"

// Session handle creation/modification of session parametr
type Session struct {
	SessionID string
	Value     phpencode.PhpSession
}

// NewSession create new instance of Session
func NewSession() *Session {
	return &Session{
		SessionID: "",
		Value:     make(phpencode.PhpSession),
	}
}

// Save session to session source
func (s *Session) Save() error {
	return nil
}
