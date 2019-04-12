package phpsessgo

// SessionIDCreator is adoptation of PHP SessionIDInterface
// Reference at https://www.php.net/manual/en/class.sessionidinterface.php
type SessionIDCreator interface {
	CreateSID() string
}
