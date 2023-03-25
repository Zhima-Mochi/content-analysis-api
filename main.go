package main

import (
	"context"
	"fmt"
	"os"
	"time"

	contentAnalysis "github.com/Zhima-Mochi/content-analysis-api/content-analysis"
)

var contentAnalysisHandler *contentAnalysis.ContentAnalysisHandler

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	contentAnalysisHandler = contentAnalysis.NewContentAnalysisHandler(apiKey)
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
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		fmt.Println("text: " + text)
		result1, err := contentAnalysisHandler.SensitiveWordsDetection(ctxWithTimeout, text)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("sensitive words dectection: " + fmt.Sprintf("%v", result1))
		// result2, err := contentAnalysisHandler.ContentClassification(ctx, text)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("content classification: " + fmt.Sprintf("%v", result2))
	}
}
