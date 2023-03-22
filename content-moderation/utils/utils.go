package utils

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
