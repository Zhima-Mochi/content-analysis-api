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

func SpamDetectionPromptGenerator(text string) string {
	return `
		## Spam Detection
		#### Input
		- text: string
		#### Output
		- is_spam: bool
		#### Example
		- text: Discount! Buy now!
		- is_spam: true
		- text: destination! Click here to claim your prize now!
		- is_spam: true
		- text: Act now and get rich quick with this amazing investment opportunity! Don't miss out on your chance to make millions!
		- is_spam: true
		- text: Urgent message: Your account has been compromised! Click here to reset your password and secure your account.
		- is_spam: true
		- text: Limited time offer: Buy one get one free on all products! Don't wait, take advantage of this incredible deal now!
		- is_spam: true
		- text: You have been selected to participate in a survey and receive a free gift card! Click here to start the survey and claim your reward.
		- is_spam: true
		- text: Reminder: Your dentist appointment is tomorrow at 2 PM. Please arrive 10 minutes early to fill out paperwork.
		- is_spam: false
		- text: Your order has been shipped and is expected to arrive on Friday. Click here to track your package.
		- is_spam: false
		- text: Thank you for your recent purchase! We hope you enjoy your new product. Please let us know if you have any questions or concerns.
		- is_spam: false
		- text: The school board meeting is scheduled for next Wednesday at 7 PM. All are welcome to attend and provide feedback on the proposed changes.
		- is_spam: false
		- text: Congratulations on your recent promotion! We are excited to see you continue to excel in your new role.
		- is_spam: false
		#### Prompt
		- text: ` + text + `
		- is_spam:
	`
}

func JudgeResult(result openai.Result) bool {
	categories := result.Categories
	scores := result.CategoryScores
	return !(categories.Hate || categories.HateThreatening || categories.SelfHarm || categories.Sexual || categories.SexualMinors || categories.Violence || scores.Sexual > 0.01)
}

// ContentSummarizationPromptGenerator generates a prompt for the Content Summarization task.
func ContentSummarizationPromptGenerator(text string) string {
	return `
		Summarizing the content with title and description
	` + text + `
	` + `
	title:
	description:
	`
}
