package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("Missing API key")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	res, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"A random joke of the day is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Choices[0].Text)
}
