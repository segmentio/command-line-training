// Simple program that will print a new log line once a second.
//
// Run it with
//
//	go run logger.go
package main

import (
	"time"

	"github.com/kevinburke/handlers"
)

func main() {
	for i := 1; true; i++ {
		handlers.Logger.Info("Another log line came in!", "svc", "command-line-training", "count", i)
		time.Sleep(1 * time.Second)
	}
}
