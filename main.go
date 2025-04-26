package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/umi-ushi/go-github-api/internal/github"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	client := github.NewGitHubClient(token)
	branch, err := github.GetDefaultBranch(client, "umi-ushi", "go-github-api")
	if err != nil {
		log.Fatal("Error fetching default branch: %v", err)
	}

	fmt.Printf("Default branch: %s", branch)
}
