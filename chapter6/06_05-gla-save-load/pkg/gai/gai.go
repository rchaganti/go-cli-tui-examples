package gai

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const WelcomeMessage = "Hello! I'm Gemini, your learning assistant. I'm here to help you learn, explore, and experiment. Ask me anything!"

type MessageContent struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

func GenerateHistory(history []MessageContent) (contents []*genai.Content) {
	for _, h := range history {
		contents = append(contents, &genai.Content{
			Parts: []genai.Part{
				genai.Text(h.Message),
			},
			Role: h.Role,
		})
	}
	return
}

func GetResponse(prompt string, history []MessageContent, apiKey string, modelName string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	model := client.GenerativeModel(modelName)
	cs := model.StartChat()
	cs.History = GenerateHistory(history)

	resp, err := cs.SendMessage(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil
}

func PrintResponse(resp *genai.GenerateContentResponse) string {
	var result string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}

	return result
}
