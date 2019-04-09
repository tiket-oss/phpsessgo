package phpsessgo

const (
	// DefaultSessionName is default session name for session config
	DefaultSessionName = "PHPSESSID"
)

// SessionConfig is adoptation of session configuration in php.ini
type SessionConfig struct {
	Name        string
	SaveHandler string
	SavePath    string
}

// DefaultSessionConfig return SessionConfig with default value
func DefaultSessionConfig() SessionConfig {
	return SessionConfig{
		Name: DefaultSessionName,
	}
}
