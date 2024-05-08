package miscellanous

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kundu-ramit/mercor_assignment/cmd/helper"
	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

func AddMiscellanousVectors() error {
	folderPath := "cmd/miscellanous/vectormap"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	var jsonData []vectorrepository.EmbeddingJSON

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(folderPath, file.Name())

			data, err := helper.ReadJSONFile(filePath)
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

func callVector(data vectorrepository.EmbeddingJSON) error {
	err := vectorrepository.NewMiscellanousRepository().Add(context.Background(), data)
	return err
}
