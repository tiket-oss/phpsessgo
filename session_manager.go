package phpsessgo

import (
	"net/http"

	"github.com/tiket-oss/phpsessgo/phpencode"
)

type SessionManager interface {
	Start(w http.ResponseWriter, r *http.Request) (session *Session, err error)
	Save(session *Session) error
	SessionName() string
	SIDCreator() SessionIDCreator
	Handler() SessionHandler
	Encoder() SessionEncoder
}

func NewSessionManager(
	sessionName string,
	sidCreator SessionIDCreator,
	handler SessionHandler,
	encoder SessionEncoder,
) SessionManager {

	return &sessionManager{
		sessionName: sessionName,
		sidCreator:  sidCreator,
		handler:     handler,
		encoder:     encoder,
	}
}

// SessionManager handle session creation/modification
type sessionManager struct {
	sessionName string
	sidCreator  SessionIDCreator
	handler     SessionHandler
	encoder     SessionEncoder
}

// Start is adoption of PHP start_session() to return current active session
func (m *sessionManager) Start(w http.ResponseWriter, r *http.Request) (session *Session, err error) {
	session = NewSession()

	var raw string
	var phpSession phpencode.PhpSession

	sessionID := m.getFromCookies(r.Cookies())

	if sessionID == "" {
		sessionID = m.sidCreator.CreateSID()
		session.SessionID = sessionID
		m.setToCookies(w, sessionID)
		return
	}

	session.SessionID = sessionID
	raw, err = m.handler.Read(sessionID)
	if err != nil {
		return
	}

	phpSession, err = m.encoder.Decode(raw)
	if err != nil {
		return
	}
	session.Value = phpSession

	return
}

// Save the session
func (m *sessionManager) Save(session *Session) error {
	sessionData, err := m.encoder.Encode(session.Value)
	if err != nil {
		return err
	}

	return m.handler.Write(session.SessionID, sessionData)
}

func (m *sessionManager) SessionName() string {
	return m.sessionName
}

func (m *sessionManager) SIDCreator() SessionIDCreator {
	return m.sidCreator
}

func (m *sessionManager) Handler() SessionHandler {
	return m.handler
}

func (m *sessionManager) Encoder() SessionEncoder {
	return m.encoder
}

func (m *sessionManager) getFromCookies(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == m.sessionName {
			return cookie.Value
		}
	}
	return ""
}

func (m *sessionManager) setToCookies(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{
		Name:     m.sessionName,
		Value:    sid,
		HttpOnly: true,
	})
}
