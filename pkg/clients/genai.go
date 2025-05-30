package clients

import (
	"bytes"
	"context"
	"codenudge/pkg/config"

	"google.golang.org/genai"
)

type GenAiClient struct {
	Config *config.GenaiClient
	Client *genai.Client
}

func NewGenAiClient(cfg *config.GenaiClient, ctx context.Context) *GenAiClient {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		panic(err)
	}

	return &GenAiClient{
		Client: client,
		Config: cfg,
	}
}

func (c *GenAiClient) UploadFile(buf *bytes.Buffer) (file *genai.File, restore func(), err error) {
	file, err = c.Client.Files.Upload(context.Background(), buf, &genai.UploadFileConfig{
		MIMEType: c.Config.FileMimeType,
	})
	restore = func() {
		c.Client.Files.Delete(context.Background(), file.Name, nil)
	}
	return
}

func (c *GenAiClient) GenerateContent(file *genai.File) (result *genai.GenerateContentResponse, err error) {
	// https://ai.google.dev/gemini-api/docs/files#upload-audio
	result, err = c.Client.Models.GenerateContent(
		context.Background(),
		c.Config.Model,
		[]*genai.Content{
			{
				Role: "user",
				Parts: []*genai.Part{
					{
						Text: c.Config.ModelPrompt,
					},
					{
						FileData: &genai.FileData{
							MIMEType: file.MIMEType,
							FileURI:  file.URI,
						},
					},
				},
			},
		},
		nil,
	)
	return
}
