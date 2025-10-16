package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":7070"
	http.HandleFunc("/hello-world", helloWorldHandler)
	log.Printf("✅ Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
