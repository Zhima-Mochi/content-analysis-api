package contentModeration

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ModerationHandler struct {
	client *openai.Client
	model  *string
}

func NewModerationHandler(client *openai.Client) *ModerationHandler {
	return &ModerationHandler{
		client: client,
	}
}

func (m *ModerationHandler) IsPass(ctx context.Context, input string) (bool, error) {
	request := openai.ModerationRequest{Input: input}
	response, err := m.client.Moderations(ctx, request)
	if err != nil {
		return false, err
	}
	categories := response.Results[0].Categories
	scores := response.Results[0].CategoryScores
	return !(categories.Hate || categories.HateThreatening || categories.SelfHarm || categories.Sexual || categories.SexualMinors || categories.Violence || scores.Sexual > 0.01), nil
}
