package helper

import (
	"encoding/json"
	"os"

	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

func CreateEmbeddingJSON(inputID, inputData string, embedding []float32) vectorrepository.EmbeddingJSON {
	return vectorrepository.EmbeddingJSON{
		InputID:   inputID,
		InputData: inputData,
		Embedding: embedding,
	}
}

func AppendJSONToFile(data vectorrepository.EmbeddingJSON, filenameprefix string, filename string) error {
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
