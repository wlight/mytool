package main

import (
	"github.com/wlight/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute errr: %v", err)
	}
}