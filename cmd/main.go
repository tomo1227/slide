package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("present", "-http", "127.0.0.1:80")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to start present: %v", err)
	}
}
