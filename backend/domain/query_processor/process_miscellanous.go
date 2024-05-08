package queryprocessor

import (
	"sort"

	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

// ProcessExperience processes the experiences based on given items.
func ProcessMiscellanous(experiences []vectorrepository.Response) Result {
	var result Result

	// Sort experiences by score in descending order
	sort.Slice(experiences, func(i, j int) bool {
		return experiences[i].Score > experiences[j].Score
	})

	for _, experience := range experiences {
		// If score is less than 0.3, ignore it
		if experience.Score < 0.3 {
			continue
		}
		// If score is greater than or equal to 0.3, set IsPresent to true
		result.IsPresent = true
		result.Responses = append(result.Responses, experience)
	}

	// If no values remain, return the value with highest score and IsPresent is false
	if len(result.Responses) == 0 && len(experiences) > 0 {
		result.Responses = []vectorrepository.Response{{Text: experiences[0].Text, Score: experiences[0].Score}}
	}

	return result
}
