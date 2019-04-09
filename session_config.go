package phpsessgo

const (
	defaultSessionName = "PHPSESSID"
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
		Name: defaultSessionName,
	}
}
