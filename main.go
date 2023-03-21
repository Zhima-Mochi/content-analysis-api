package main

import (
	"context"
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
		"哇操",
		"我不要乖乖聽話",
		"每天早起做愛心便當",
	}
	for _, text := range texts {
		log.Println(text)
		result, err := contentModerationHandler.SensitiveWordsDetection(ctx, text)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(result)
	}
}
