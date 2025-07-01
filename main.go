package main

import (
	"api-test/cmd"
	"api-test/config"
	"log"
)

func main() {
	if err := config.LoadConfig("config/config.yml"); err != nil {
		log.Fatal("error:", err)
	}

	cmd.Start()
}
