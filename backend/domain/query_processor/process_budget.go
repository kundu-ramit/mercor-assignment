package queryprocessor

import (
	"sort"

	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

// Result represents the result of the ProcessBudget function.
type Result struct {
	Responses []vectorrepository.Response
	IsPresent bool
}

// ProcessBudget processes the budget based on given responses.
func ProcessBudget(responses []vectorrepository.Response) Result {
	var result Result

	// Sort responses by score in descending order
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Score > responses[j].Score
	})

	for _, response := range responses {
		// If score is less than 0.3, ignore it
		if response.Score < 0.3 {
			continue
		}
		// If score is greater than or equal to 0.3, set IsPresent to true
		result.IsPresent = true
		result.Responses = append(result.Responses, response)
	}

	// If no values remain, return the value with highest score and IsPresent is false
	if len(result.Responses) == 0 && len(responses) > 0 {
		result.Responses = []vectorrepository.Response{{Text: responses[0].Text, Score: responses[0].Score}}
	}

	return result
}
