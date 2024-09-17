package gai

import (
	"context"
	"fmt"
	"log"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func ListModels(apiKey string) []*genai.ModelInfo {
	ctx := context.Background()
	models := make([]*genai.ModelInfo, 0)
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		models = append(models, m)
	}

	return models
}

func GetModel(apiKey, name string) genai.ModelInfo {
	var model genai.ModelInfo
	models := ListModels(apiKey)
	for _, m := range models {
		if m.Name == name {
			model = *m
		}
	}
	return model
}

func GenerateFromText(apikey, prompt string) *genai.GenerateContentResponse {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func GenerateFromImage(apikey, fromImage, prompt string) *genai.GenerateContentResponse {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	img, err := os.ReadFile(fromImage)
	if err != nil {
		log.Fatal(err)
	}

	req := []genai.Part{
		genai.ImageData("png", img),
		genai.Text(prompt),
	}
	resp, err := model.GenerateContent(ctx, req...)

	if err != nil {
		log.Fatal(err)
	}

	return resp
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
