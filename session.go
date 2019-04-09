package phpsessgo

// Session handle creation/modification of session parametr
type Session interface {
	SessionID() string
	Get(key string) interface{}
	Put(key string, val interface{})
}

// NewSession create new instance of Session
func NewSession(sid string) Session {
	return &session{
		sid: sid,
	}
}

type session struct {
	sid string
}

func (s *session) SessionID() string {
	return s.sid
}

func (s *session) Get(key string) interface{} {
	return nil
}

func (s *session) Put(key string, val interface{}) {

}
