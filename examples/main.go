package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ahmedthabet/serper-go/serper"
)

func main() {
	apiKey := os.Getenv("SERPER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: SERPER_API_KEY is not set")
		os.Exit(1)
	}

	// Initialize client with functional options
	client := serper.NewClient(
		apiKey,
		serper.WithTimeout(10*time.Second),
	)

	ctx := context.Background()

	// 1. Organic Search with Constants
	fmt.Println("Searching for 'Go programming' in the US...")
	req := &serper.Request{
		Q:  "Go programming",
		Gl: serper.GLUnitedStates,
		Hl: serper.HLEnglish,
	}

	search, err := client.Search(ctx, req)
	if err != nil {
		handleError(err)
	} else {
		fmt.Printf("Results (Credits: %d):\n", search.Credits)
		for _, r := range search.Organic {
			fmt.Printf("- %s (%s)\n", r.Title, r.Link)
		}
	}

	// 2. Image Search
	fmt.Println("\nSearching for images...")
	images, err := client.Images(ctx, &serper.Request{Q: "Gopher mascot"})
	if err != nil {
		handleError(err)
	} else {
		for i, r := range images.Images {
			if i >= 3 {
				break
			}
			fmt.Printf("- %s\n", r.ImageUrl)
		}
	}
}

func handleError(err error) {
	var apiErr *serper.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("Serper API Error: %s (Status: %d)\n", apiErr.Message, apiErr.StatusCode)
		return
	}
	fmt.Printf("Unexpected Error: %v\n", err)
}
