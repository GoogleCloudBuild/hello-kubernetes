// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
