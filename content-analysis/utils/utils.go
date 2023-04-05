package utils

import "github.com/sashabaranov/go-openai"

// SensitiveWordsDetectionPromptGenerator generates a prompt for the Sensitive Words Detection task.
func SensitiveWordsDetectionPromptGenerator(text string) string {
	return `
		## Judging the sensitivity of the text
		#### Input
		- text: string
		#### Output
		- is_sensitive: bool
		#### Example
		- text: fuck
		- is_sensitive: true
		#### Prompt
		- text: ` + text + `
		- is_sensitive:
	`
}

// ContentClassificationPromptGenerator generates a prompt for the Content Classification task.
func ContentClassificationPromptGenerator(text string) string {
	return `
		## Classifying the content only one category (unknown category is not allowed)
		#### Input
		- text: string
		#### Output
		- category: string
		#### Example
		- text: I want to be a good person
		- category: positive
		#### Prompt
		- text: ` + text + `
		- category:
	`
}

func JudgeResult(result openai.Result) bool {
	categories := result.Categories
	scores := result.CategoryScores
	return !(categories.Hate || categories.HateThreatening || categories.SelfHarm || categories.Sexual || categories.SexualMinors || categories.Violence || scores.Sexual > 0.01)
}

// ContentSummarizationPromptGenerator generates a prompt for the Content Summarization task.
func ContentSummarizationPromptGenerator(text string) string {
	return `Summarizing the content with title in the first line and description in the next line:
` + text
}
