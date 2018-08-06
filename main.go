// A very simple Go application.
package main

import (
	"fmt"
	"log"
	"time"
)

// Greet returns a pleasant greeting.
func Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	greeting := Greet("change-me")
	log.Printf(greeting)
	for {
		// Don't exit, otherwise Kubernetes thinks we crashed.
		time.Sleep(10 * time.Second)
	}
}
