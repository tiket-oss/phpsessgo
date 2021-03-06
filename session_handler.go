package phpsessgo

// SessionHandler is adoption of PHP SessionHandlerInterface
// For more reference: https://www.php.net/manual/en/class.sessionhandlerinterface.php
type SessionHandler interface {
	Close()
	// Gc(maxLifeTime int) int
	Read(sessionID string) (string, error)
	Write(sessionID, sessionData string) error
}
