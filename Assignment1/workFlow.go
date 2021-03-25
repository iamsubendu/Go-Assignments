package main

import (
	"fmt"
	"time"
)

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}

func main() {
	started := time.Now()
	foods := []string{"aata", "ghee", "fire", "hotpot"}
	for _, f := range foods {
		cook(f)
	}
	fmt.Printf("done in %v\n", time.Since(started))
}