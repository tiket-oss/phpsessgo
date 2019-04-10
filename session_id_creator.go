package phpsessgo

// SessionIDCreator is adoptation of PHP SessionIDInterface
// Reference at https://www.php.net/manual/en/class.sessionidinterface.php
type SessionIDCreator interface {
	CreateSID() string
}

type sessionIDCreator struct {
	SessionIDCreator
}

func (g *sessionIDCreator) CreateSID() string {
	// INFO: Reference at https://stackoverflow.com/questions/12240922/what-is-the-length-of-a-php-session-id-string
	return "thedroppedcookiehasgoldinit"
}
