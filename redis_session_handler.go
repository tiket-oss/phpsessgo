package phpsessgo

type redisSessionHandler struct {
	config SessionConfig
}

func (h *redisSessionHandler) Open(savePath, sessionName string) bool {
	return false
}

func (h *redisSessionHandler) Close() bool {
	return false
}

func (h *redisSessionHandler) Gc(maxLifeTime int) int {
	return -1
}

func (h *redisSessionHandler) Read(sessionID string) string {
	return ""
}

func (h *redisSessionHandler) Write(sessionID string, sessionData interface{}) string {
	return ""
}
