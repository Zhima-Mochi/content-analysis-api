package contentModeration

import (
	"context"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type ContentModerationHandler struct {
	client      *openai.Client
	model       string
	temperature float32
}

func NewContentModerationHandler(client *openai.Client) *ContentModerationHandler {
	return &ContentModerationHandler{
		client:      client,
		model:       openai.GPT3Dot5Turbo0301,
		temperature: 0.5,
	}
}

func (h *ContentModerationHandler) SetModel(model string) {
	h.model = model
}

func (h *ContentModerationHandler) GetModel() string {
	return h.model
}

func (h *ContentModerationHandler) SetTemperature(temperature float32) {
	h.temperature = temperature
}

func (h *ContentModerationHandler) GetTemperature() float32 {
	return h.temperature
}

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

// generateSensitiveWordsDetectionPrompt generates a prompt for the Sensitive Words Detection task.
func generateSensitiveWordsDetectionPrompt(text string) string {
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

func (h *ContentModerationHandler) SensitiveWordsDetection(ctx context.Context, text string) (bool, error) {
	prompt := generateSensitiveWordsDetectionPrompt(text)
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

func (h *ContentModerationHandler) ContentClassification(ctx context.Context, text string) (string, error) {
	return "", nil
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
