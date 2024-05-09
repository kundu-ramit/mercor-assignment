package miscellanous

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kundu-ramit/mercor_assignment/cmd/helper"
	"github.com/kundu-ramit/mercor_assignment/domain/openai"
)

type Miscellanous struct {
	ID    string `gorm:"column:id"`
	Value string `gorm:"column:value"`
}

func FetchMiscellanousVectorOpenAi() {
	filename := "cmd/miscellanous/miscellanous_list.json"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var miscellanous []Miscellanous
	err = json.NewDecoder(file).Decode(&miscellanous)
	if err != nil {
		panic(err)
	}

	for _, miscellanous := range miscellanous {
		err := processMiscellanous(miscellanous.Value, miscellanous.ID)
		if err != nil {
			fmt.Printf("Error processing skill '%s': %v\n", miscellanous.ID, err)
			continue
		}
		fmt.Printf("Skill '%s' processed successfully\n", miscellanous.ID)
	}
}

func processMiscellanous(value, id string) error {
	filename := "cmd/miscellanous/vectormap/"

	embedding, err := openai.NewOpenAiProcessor().GetEmbeddingVector(value)
	if err != nil {
		panic(err)
	}

	embeddingJSON := helper.CreateEmbeddingJSON(id, id, embedding)

	err = helper.AppendJSONToFile(embeddingJSON, filename, id)
	if err != nil {
		return err
	}

	return nil
}
