package phpsessgo

// SessionHandler is adoption of PHP SessionHandlerInterface
// For more reference: https://www.php.net/manual/en/class.sessionhandlerinterface.php
var database map[string]string

func init() {
	database = make(map[string]string)
}

func Read(sessionID string) (data string, err error) {
	if value, ok := database[sessionID]; ok {
		return value, nil
	}
	return "", nil
}

func Write(sessionID string, sessionData string) error {
	database[sessionID] = sessionData
	return nil
}
