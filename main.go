package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "failed to load .env: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Println("gh-project-promoter")
}
