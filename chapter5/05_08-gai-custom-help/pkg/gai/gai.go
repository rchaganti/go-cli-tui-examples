package gai

import (
	"context"
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func ListModels(apiKey string) ([]*genai.ModelInfo, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	ctx := context.Background()
	models := make([]*genai.ModelInfo, 0)
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		models = append(models, m)
	}

	return models, nil
}

func GetModel(apiKey, name string) (genai.ModelInfo, error) {
	if apiKey == "" {
		return genai.ModelInfo{}, fmt.Errorf("API key is required")
	}
	var model genai.ModelInfo
	models, err := ListModels(apiKey)
	if err != nil {
		return genai.ModelInfo{}, err
	}
	for _, m := range models {
		if m.Name == name {
			model = *m
		}
	}
	return model, nil
}

func GenerateFromText(apikey, prompt string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	if apikey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GenerateFromImage(apikey, fromImage, prompt string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	if apikey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	img, err := os.ReadFile(fromImage)
	if err != nil {
		return nil, err
	}

	req := []genai.Part{
		genai.ImageData("png", img),
		genai.Text(prompt),
	}
	resp, err := model.GenerateContent(ctx, req...)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func PrintResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			stringResp := fmt.Sprintf("%s", part)
			result := markdown.Render(stringResp, 80, 6)
			fmt.Println(string(result))
		}
	}
}
