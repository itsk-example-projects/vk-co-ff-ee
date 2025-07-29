package web

import (
	"log"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/", indexHandler)
	//log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
