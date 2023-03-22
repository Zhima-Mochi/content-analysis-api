package main

import (
	"context"
	"fmt"
	"log"
	"os"

	contentModeration "github.com/Zhima-Mochi/content-moderation-api/content-moderation"
	"github.com/sashabaranov/go-openai"
)

var contentModerationHandler *contentModeration.ContentModerationHandler

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	contentModerationHandler = contentModeration.NewContentModerationHandler(openai.NewClient(apiKey))
}

func main() {
	ctx := context.Background()
	texts := []string{
		"我不要听话",
		"you are so stupid",
		"每天早起做愛心便當",
		"この野郎！",
	}
	for _, text := range texts {
		fmt.Println(text)
		result, err := contentModerationHandler.SensitiveWordsDetection(ctx, text)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(result)
	}
}
