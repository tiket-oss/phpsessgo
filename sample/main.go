package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	sessionName = "PHPSESSID" // same with session.name in phpinfo()
)

var (
	name = flag.String("name", "echo-01", "")
	addr = flag.String("address", ":8181", "echo server address")
)

func main() {
	flag.Parse()

	fmt.Printf(`%s listen and serve at %s`, *name, *addr)

	// server
	http.HandleFunc("/", handleFunc)

	err := http.ListenAndServe(*addr, nil)
	fatalIfError(err)
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
