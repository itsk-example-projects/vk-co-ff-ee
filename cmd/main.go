package main

import (
	"vk-co-ff-ee/internal/config"
	"vk-co-ff-ee/internal/logger"
	"vk-co-ff-ee/internal/web"
)

func main() {
	config.LoadConfig()
	logger.SetupLogging()
	web.StartWebServer()
}
