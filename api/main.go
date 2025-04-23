package main

import (
	"log"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	appDir := filepath.Dir(ex)
	envPath := filepath.Join(appDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Error loading .env from %s: %v", envPath, err)
	}

	log.Fatal(startServer(":1345"))
}
