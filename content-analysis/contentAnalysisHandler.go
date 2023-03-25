package contentAnalysis

import (
	"context"
	"strings"

	"github.com/Zhima-Mochi/content-analysis-api/content-analysis/utils"
	"github.com/sashabaranov/go-openai"
)

type ContentAnalysisHandler struct {
	client                                 *openai.Client
	userLanguage                           string
	model                                  string
	temperature                            float32
	sensitiveWordsDetectionPromptGenerator func(text string) string
	contentClassificationGenerator         func(text string) string
	spamDetectionPromptGenerator           func(text string) string
	contentSummarizaioinGenerator          func(text string) string
	moderationHandler                      *ModerationHandler
}

func NewContentAnalysisHandler(apiKey string) *ContentAnalysisHandler {
	client := openai.NewClient(apiKey)
	return &ContentAnalysisHandler{
		client:                                 client,
		userLanguage:                           "english",
		model:                                  openai.GPT3Dot5Turbo0301,
		temperature:                            0.5,
		sensitiveWordsDetectionPromptGenerator: utils.SensitiveWordsDetectionPromptGenerator,
		contentClassificationGenerator:         utils.ContentClassificationPromptGenerator,
		spamDetectionPromptGenerator:           utils.SpamDetectionPromptGenerator,
		contentSummarizaioinGenerator:          utils.ContentSummarizationPromptGenerator,
		moderationHandler:                      NewModerationHandler(client),
	}
}

// SetUserLanguage sets the user language of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) SetUserLanguage(language string) {
	h.userLanguage = language
}

// GetUserLanguage returns the user language of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) GetUserLanguage() string {
	return h.userLanguage
}

func (h *ContentAnalysisHandler) SetModerationHandlerJudgeResult(judgeResult func(openai.Result) bool) {
	h.moderationHandler.SetJudgeResult(judgeResult)
}

// SetModel sets the model of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) SetModel(model string) {
	h.model = model
}

// GetModel returns the model of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) GetModel() string {
	return h.model
}

// SetTemperature sets the temperature of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) SetTemperature(temperature float32) {
	h.temperature = temperature
}

// GetTemperature returns the temperature of the ContentAnalysisHandler.
func (h *ContentAnalysisHandler) GetTemperature() float32 {
	return h.temperature
}

// getCompletionWithContent is a helper function that returns the completion with the content.
func (h *ContentAnalysisHandler) getCompletionWithContent(ctx context.Context, prompt string) (string, error) {
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

// getCompletionWithMessages is a helper function that returns the completion with the messages.
func (h *ContentAnalysisHandler) getCompletionWithMessages(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:       h.model,
		Messages:    messages,
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
func (h *ContentAnalysisHandler) SetSensitiveWordsDetectionPromptGenerator(generator func(text string) string) {
	if generator == nil {
		return
	}
	h.sensitiveWordsDetectionPromptGenerator = generator
}

// SensitiveWordsDetection detects sensitive words in the text.
func (h *ContentAnalysisHandler) SensitiveWordsDetection(ctx context.Context, text string) (bool, error) {
	answer := "true"
	var err error
	prompt := h.sensitiveWordsDetectionPromptGenerator(text)
	answer, err = h.getCompletionWithContent(ctx, prompt)
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
func (h *ContentAnalysisHandler) SetContentClassificationGenerator(generator func(text string) string) {
	if generator == nil {
		return
	}
	h.contentClassificationGenerator = generator
}

// ContentClassification classifies the text.
func (h *ContentAnalysisHandler) ContentClassification(ctx context.Context, text string) (string, error) {
	prompt := h.contentClassificationGenerator(text)
	answer, err := h.getCompletionWithContent(ctx, prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}

func (h *ContentAnalysisHandler) ContentSimilarityDetection(ctx context.Context, text1, text2 string) (bool, error) {
	return false, nil
}

func (h *ContentAnalysisHandler) ContentSummarization(ctx context.Context, text string) (string, error) {
	prompt := h.contentSummarizaioinGenerator(text)
	prompt += `
	(please respond to me in ` + h.userLanguage + `)
	`
	answer, err := h.getCompletionWithContent(ctx, prompt)
	if err != nil {
		return "", err
	}

	return answer, nil
}
