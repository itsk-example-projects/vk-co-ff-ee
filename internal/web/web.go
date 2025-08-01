package web

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"vk-co-ff-ee/internal/config"
)

func StartWebServer() {
	http.HandleFunc("/", indexHandler)
	//log.Println("Starting server on :8080")
	if err := http.ListenAndServe(viper.GetString(config.Port), nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
