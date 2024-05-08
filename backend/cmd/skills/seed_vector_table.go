package skills

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

func AddSkillVectors() error {
	folderPath := "cmd/skills/vectormap"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var jsonData []vectorrepository.EmbeddingJSON

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(folderPath, file.Name())

			data, err := readJSONFile(filePath)
			if err != nil {
				fmt.Printf("Error reading JSON file '%s': %v\n", filePath, err)
				continue
			}

			jsonData = append(jsonData, data)
		}
	}

	for _, data := range jsonData {
		err := callVector(data)
		if err != nil {
			fmt.Printf("Error processing data '%s': %v\n", data.InputID, err)
			continue
		}
		fmt.Printf("Data '%s' processed successfully\n", data.InputID)
	}

	return nil
}

func readJSONFile(filePath string) (vectorrepository.EmbeddingJSON, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return vectorrepository.EmbeddingJSON{}, err
	}
	defer file.Close()

	var data vectorrepository.EmbeddingJSON
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return vectorrepository.EmbeddingJSON{}, err
	}

	return data, nil
}

func callVector(data vectorrepository.EmbeddingJSON) error {
	err := vectorrepository.NewSkillRepository().Add(context.Background(), data)
	return err
}
