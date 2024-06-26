package budget

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kundu-ramit/mercor_assignment/cmd/helper"
	"github.com/kundu-ramit/mercor_assignment/domain/openai"
)

type Budget struct {
	ID    string `gorm:"column:id"`
	Value string `gorm:"column:value"`
}

func FetchBudgetVectorOpenAi() {
	filename := "cmd/budget/budget_list.json"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var budget []Budget
	err = json.NewDecoder(file).Decode(&budget)
	if err != nil {
		panic(err)
	}

	for _, budget := range budget {
		err := processBudget(budget.Value, budget.ID)
		if err != nil {
			fmt.Printf("Error processing skill '%s': %v\n", budget.ID, err)
			continue
		}
		fmt.Printf("Skill '%s' %s processed successfully\n", budget.ID, budget.Value)
	}
}

func processBudget(value, id string) error {
	filename := "cmd/budget/vectormap/"

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
