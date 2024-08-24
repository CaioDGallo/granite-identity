package main

import (
	"log"
	"os/exec"

	"github.com/CaioDGallo/granite-identity/internal/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cmd := exec.Command("tern", "migrate", "--migrations", "./db/migrations", "--config", "./db/tern.conf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running tern migrate: %v\nOutput: %s", err, string(output))
		panic(err)
	}

	log.Printf("Command output: %s", string(output))
}
