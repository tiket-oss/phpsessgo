package main

import (
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	sessionID := sessionID(r.Cookies())
	if sessionID == "" {
		sessionID = generateSessionID()
		http.SetCookie(w, &http.Cookie{
			Name:  sessionName,
			Value: sessionID,
		})
	}

	w.Write([]byte(sessionID))
}

func generateSessionID() string {
	// https://stackoverflow.com/questions/12240922/what-is-the-length-of-a-php-session-id-string
	return "thedroppedcookiehasgoldinit"
}

func sessionID(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == sessionName {
			return cookie.Value
		}
	}
	return ""

}
