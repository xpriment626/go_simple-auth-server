package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-server-auth/internal"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	fmt.Printf("Server running at port: %s\n", CONN_PORT)
	http.HandleFunc("/", internal.BasicAuth(internal.Message, "Please enter your username and password"))
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server:", err)
		return
	}
}
