package phpsessgo

// SessionHandler is adoption of PHP SessionHandlerInterface
// For more reference: https://www.php.net/manual/en/class.sessionhandlerinterface.php
type SessionHandler interface {
	Open(savePath, sessionName string) bool
	Close() bool
	Gc(maxLifeTime int) int
	Read(sessionID string) interface{}
	Write(sessionID string, sessionData interface{}) bool
}
