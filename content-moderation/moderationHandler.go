package contentModeration

import (
	"context"

	"github.com/Zhima-Mochi/content-moderation-api/content-moderation/utils"
	"github.com/sashabaranov/go-openai"
)

type ModerationHandler struct {
	client *openai.Client
	model  *string
	judge  func(openai.ModerationResponse) bool
}

func NewModerationHandler(client *openai.Client) *ModerationHandler {
	return &ModerationHandler{
		client: client,
		judge:  utils.Judge,
	}
}

func (m *ModerationHandler) IsPass(ctx context.Context, input string) (bool, error) {
	request := openai.ModerationRequest{Input: input}
	response, err := m.client.Moderations(ctx, request)
	if err != nil {
		return false, err
	}
	return m.judge(response), nil
}

// SetJudge sets the judge of the ModerationHandler.
func (m *ModerationHandler) SetJudge(judge func(openai.ModerationResponse) bool) {
	m.judge = judge
}
