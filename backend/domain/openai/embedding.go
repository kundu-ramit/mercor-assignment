package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Index     int       `json:"index"`
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type OpenAiProcessor interface {
	GetEmbeddingVector(inputData string) ([]float32, error)
}

func NewOpenAiProcessor() openAiProcessor {
	return openAiProcessor{}
}

type openAiProcessor struct {
}

func (oai openAiProcessor) GetEmbeddingVector(inputData string) ([]float32, error) {
	bToken := os.Getenv("GPT_KEY")
	url := "https://api.openai.com/v1/embeddings"

	data := map[string]interface{}{
		"input": inputData,
		"model": "text-embedding-3-small",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body bytes.Buffer
	_, err = io.Copy(&body, resp.Body)
	if err != nil {
		return nil, err
	}

	var embeddingResponse EmbeddingResponse
	err = json.Unmarshal(body.Bytes(), &embeddingResponse)
	if err != nil {
		return nil, err
	}

	if len(embeddingResponse.Data) > 0 {
		return embeddingResponse.Data[0].Embedding, nil
	}

	return nil, fmt.Errorf("no embedding data found in the response")
}
