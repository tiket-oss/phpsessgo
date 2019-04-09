package phpsessgo

// SessionIDGenerator is adoptation of PHP SessionIDInterface
// Reference at https://www.php.net/manual/en/class.sessionidinterface.php
type SessionIDGenerator interface {
	CreateSID() string
}

// NewSessionIDGenerator return instance of SessionIDGenerator
func NewSessionIDGenerator() SessionIDGenerator {
	return &sessionIDGenerator{}
}

type sessionIDGenerator struct {
}

func (g *sessionIDGenerator) CreateSID() string {
	// INFO: Reference at https://stackoverflow.com/questions/12240922/what-is-the-length-of-a-php-session-id-string
	return "thedroppedcookiehasgoldinit"
}
