package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/imantung/phpsessgo"
)

var (
	addr = flag.String("address", ":8181", "echo server address")
)

var sessionManager *phpsessgo.SessionManager

func main() {
	flag.Parse()
	fmt.Printf("Listen and serve at %s\n", *addr)

	var err error

	// initiatte SessionManager
	sessionConfig := phpsessgo.DefaultSessionConfig()
	sessionManager, err = phpsessgo.NewSessionManager(sessionConfig)
	fatalIfError(err)

	// server
	http.HandleFunc("/", handleFunc)

	err = http.ListenAndServe(*addr, nil)
	fatalIfError(err)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	session, err := sessionManager.Start(w, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	if session != nil {
		w.Write([]byte(session.SessionID()))
		w.Write([]byte("\n"))
		w.Write([]byte(session.Value["spike01"].(string)))
	}

}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
