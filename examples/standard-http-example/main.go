package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/tiket-oss/phpsessgo"
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
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	sessionManager = phpsessgo.NewRedisSessionManager(client)

	// server
	http.HandleFunc("/", handleFunc)

	err = http.ListenAndServe(*addr, nil)
	fatalIfError(err)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	// PHP: session_start();
	session, err := sessionManager.Start(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer sessionManager.Save(session)

	// PHP: $_SESSION["hello"] = "world";
	session.Value["hello"] = "world"

	// PHP: session_Id();
	w.Write([]byte(session.SessionID))
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
