package phpsessgo

// SessionIDCreator is adoptation of PHP SessionIDInterface
// Reference at https://www.php.net/manual/en/class.sessionidinterface.php
type SessionIDCreator interface {
	CreateSID() string
}

// NewSessionIDCreator return instance of SessionIDGenerator
func NewSessionIDCreator() SessionIDCreator {
	return &sessionIDCreator{}
}

type sessionIDCreator struct {
}

func (g *sessionIDCreator) CreateSID() string {
	// INFO: Reference at https://stackoverflow.com/questions/12240922/what-is-the-length-of-a-php-session-id-string
	return "thedroppedcookiehasgoldinit"
}
