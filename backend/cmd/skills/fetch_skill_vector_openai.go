package skills

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kundu-ramit/mercor_assignment/domain/openai"
	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

func createEmbeddingJSON(inputID, inputData string, embedding []float32) vectorrepository.EmbeddingJSON {
	return vectorrepository.EmbeddingJSON{
		InputID:   inputID,
		InputData: inputData,
		Embedding: embedding,
	}
}

func appendJSONToFile(data vectorrepository.EmbeddingJSON, filenameprefix string, filename string) error {
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

	embedding, err := openai.GetEmbeddingVector(inputData, bToken)
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
