package skills

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

type EmbeddingJSON struct {
	InputID   string    `json:"inputID"`
	InputData string    `json:"inputData"`
	Embedding []float32 `json:"embedding"`
}

func makeCurlRequest(inputData string, bToken string) ([]float32, error) {
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

func createEmbeddingJSON(inputID, inputData string, embedding []float32) EmbeddingJSON {
	return EmbeddingJSON{
		InputID:   inputID,
		InputData: inputData,
		Embedding: embedding,
	}
}

func appendJSONToFile(data EmbeddingJSON, filenameprefix string, filename string) error {
	filePath := filenameprefix + filename + ".json"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonData))
	if err != nil {
		return err
	}

	return nil
}

func FetchSkillVectorOpenAi() {
	filename := "cmd/skills/skill_list.json"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var skills []Skill
	err = json.NewDecoder(file).Decode(&skills)
	if err != nil {
		panic(err)
	}

	for _, skill := range skills {
		err := processSkill(skill.CustomText, skill.SkillID, skill.SkillName)
		if err != nil {
			fmt.Printf("Error processing skill '%s': %v\n", skill.SkillName, err)
			continue
		}
		fmt.Printf("Skill '%s' processed successfully\n", skill.SkillName)
	}
}

func processSkill(inputData, inputID string, inputName string) error {
	bToken := os.Getenv("GPT_KEY")
	filename := "cmd/skills/vectormap/"

	embedding, err := makeCurlRequest(inputData, bToken)
	if err != nil {
		panic(err)
	}

	embeddingJSON := createEmbeddingJSON(inputID, inputData, embedding)

	err = appendJSONToFile(embeddingJSON, filename, inputName)
	if err != nil {
		return err
	}

	return nil
}
