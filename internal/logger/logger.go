package logger

import (
	"io"
	"log"
	"os"
)

func SetupLogging() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open/create log file: %v", err)
	}
	// No closing here, as the file should remain open for the lifetime of the application and when a process terminates via "execve" or normal exit, the kernel closes all open file descriptors automatically
	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
