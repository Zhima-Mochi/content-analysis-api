package main

import (
	"context"
	"fmt"
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
		"幹你娘",
	}
	for _, text := range texts {
		fmt.Println("text: " + text)
		result1, err := contentModerationHandler.SensitiveWordsDetection(ctx, text)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("sensitive words dectection: " + fmt.Sprintf("%v", result1))
		// result2, err := contentModerationHandler.ContentClassification(ctx, text)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("content classification: " + fmt.Sprintf("%v", result2))
	}
}
