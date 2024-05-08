package experience

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kundu-ramit/mercor_assignment/cmd/helper"
	"github.com/kundu-ramit/mercor_assignment/domain/openai"
)

type Experience struct {
	ID    string `gorm:"column:id"`
	Value string `gorm:"column:value"`
}

func FetchExperienceVectorOpenAi() {
	filename := "cmd/experience/experience_list.json"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var experience []Experience
	err = json.NewDecoder(file).Decode(&experience)
	if err != nil {
		panic(err)
	}

	for _, experience := range experience {
		err := processExperience(experience.Value, experience.ID)
		if err != nil {
			fmt.Printf("Error processing skill '%s': %v\n", experience.ID, err)
			continue
		}
		fmt.Printf("Skill '%s' processed successfully\n", experience.ID)
	}
}

func processExperience(value, id string) error {
	bToken := os.Getenv("GPT_KEY")
	filename := "cmd/experience/vectormap/"

	embedding, err := openai.GetEmbeddingVector(value, bToken)
	if err != nil {
		panic(err)
	}

	embeddingJSON := helper.CreateEmbeddingJSON(id, "", embedding)

	err = helper.AppendJSONToFile(embeddingJSON, filename, "")
	if err != nil {
		return err
	}

	return nil
}
