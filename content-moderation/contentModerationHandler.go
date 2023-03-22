package contentModeration

import (
	"context"
	"strings"

	"github.com/Zhima-Mochi/content-moderation-api/content-moderation/utils"
	"github.com/sashabaranov/go-openai"
)

type ContentModerationHandler struct {
	client                                 *openai.Client
	model                                  string
	temperature                            float32
	sensitiveWordsDetectionPromptGenerator func(text string) string
	contentClassificationGenerator         func(text string) string
	spamDetectionPromptGenerator           func(text string) string
}

func NewContentModerationHandler(client *openai.Client) *ContentModerationHandler {
	return &ContentModerationHandler{
		client:                                 client,
		model:                                  openai.GPT3Dot5Turbo0301,
		temperature:                            0.5,
		sensitiveWordsDetectionPromptGenerator: utils.SensitiveWordsDetectionPromptGenerator,
		contentClassificationGenerator:         utils.ContentClassificationPromptGenerator,
		spamDetectionPromptGenerator:           utils.SpamDetectionPromptGenerator,
	}
}

// SetModel sets the model of the ContentModerationHandler.
func (h *ContentModerationHandler) SetModel(model string) {
	h.model = model
}

// GetModel returns the model of the ContentModerationHandler.
func (h *ContentModerationHandler) GetModel() string {
	return h.model
}

// SetTemperature sets the temperature of the ContentModerationHandler.
func (h *ContentModerationHandler) SetTemperature(temperature float32) {
	h.temperature = temperature
}

// GetTemperature returns the temperature of the ContentModerationHandler.
func (h *ContentModerationHandler) GetTemperature() float32 {
	return h.temperature
}

// contentGeneration generates content based on the prompt.
func (h *ContentModerationHandler) contentGeneration(ctx context.Context, prompt string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: h.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: h.temperature,
	}

	resp, err := h.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

// SetSensitiveWordsDetectionPromptGenerator sets the prompt generator for the Sensitive Words Detection task.
// The default prompt generator is utils.SensitiveWordsDetectionPromptGenerator.
// The prompt generator should be a function that takes a string as input and returns a string as output.
// The input string is the text to be judged.
// The output string is the prompt.
// For example:
//
//	func SensitiveWordsDetectionPromptGenerator(text string) string {
//		return `
//			## Judging the sensitivity of the text
//			#### Input
//			- text: string
//			#### Output
//			- is_sensitive: bool
//			#### Example
//			- text: fuck
//			- is_sensitive: true
//			#### Prompt
//			- text: ` + text + `
//			- is_sensitive:
//		`
//	}
func (h *ContentModerationHandler) SetSensitiveWordsDetectionPromptGenerator(generator func(text string) string) {
	h.sensitiveWordsDetectionPromptGenerator = generator
}

// SensitiveWordsDetection detects sensitive words in the text.
func (h *ContentModerationHandler) SensitiveWordsDetection(ctx context.Context, text string) (bool, error) {
	prompt := h.sensitiveWordsDetectionPromptGenerator(text)
	answer, err := h.contentGeneration(ctx, prompt)
	if err != nil {
		return false, err
	}
	answer = strings.ToLower(answer)

	if strings.Contains(answer, "true") {
		return true, nil
	} else if strings.Contains(answer, "false") {
		return false, nil
	}

	return false, ErrInvalidAnswer
}

// SetContentClassificationGenerator sets the prompt generator for the Content Classification task.
// The default prompt generator is utils.ContentClassificationPromptGenerator.
// The prompt generator should be a function that takes a string as input and returns a string as output.
// The input string is the text to be judged.
// The output string is the prompt.
// For example:
//
//	func ContentClassificationPromptGenerator(text string) string {
//		return `
//			## Judging the classification of the text
//			#### Input
//			- text: string
//			#### Output
//			- classification: string
//			#### Example
//			- text: I want to be a doctor.
//			- classification: career
//			#### Prompt
//			- text: ` + text + `
//			- classification:
//		`
//	}
func (h *ContentModerationHandler) SetContentClassificationGenerator(generator func(text string) string) {
	h.contentClassificationGenerator = generator
}

// ContentClassification classifies the text.
func (h *ContentModerationHandler) ContentClassification(ctx context.Context, text string) (string, error) {
	prompt := h.contentClassificationGenerator(text)
	answer, err := h.contentGeneration(ctx, prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}

// SetSpamDetectionPromptGenerator sets the prompt generator for the Spam Detection task.
func (h *ContentModerationHandler) SetSpamDetectionPromptGenerator(generator func(text string) string) {
	h.spamDetectionPromptGenerator = generator
}

func (h *ContentModerationHandler) SpamDetection(ctx context.Context, text string) (bool, error) {
	return false, nil
}

func (h *ContentModerationHandler) ContentSimilarityDetection(ctx context.Context, text1, text2 string) (bool, error) {
	return false, nil
}

func (h *ContentModerationHandler) ContentSummarization(ctx context.Context, text string) (string, error) {
	return "", nil
}
