package contentAnalysis

import (
	"context"

	"github.com/Zhima-Mochi/content-analysis-api/utils"
	"github.com/sashabaranov/go-openai"
)

type ModerationHandler struct {
	client      *openai.Client
	model       *string
	judgeResult func(openai.Result) bool
}

func NewModerationHandler(client *openai.Client) *ModerationHandler {
	return &ModerationHandler{
		client:      client,
		judgeResult: utils.JudgeResult,
	}
}

func (m *ModerationHandler) IsPass(ctx context.Context, input string) (bool, error) {
	request := openai.ModerationRequest{Input: input}
	response, err := m.client.Moderations(ctx, request)
	if err != nil {
		return false, err
	}
	return m.judgeResult(response.Results[0]), nil
}

// SetJudgeResult sets the judgeResult of the ModerationHandler.
func (m *ModerationHandler) SetJudgeResult(judgeResult func(openai.Result) bool) {
	m.judgeResult = judgeResult
}
