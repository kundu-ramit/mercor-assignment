package queryprocessor

import (
	"context"

	mysqlrepository "github.com/kundu-ramit/mercor_assignment/domain/mysql_repository"
	"github.com/kundu-ramit/mercor_assignment/domain/openai"
	vectorrepository "github.com/kundu-ramit/mercor_assignment/domain/vector_repository"
)

type ProcessorResponse struct {
	Skills       Result
	Budget       Result
	Experience   Result
	Miscellanous Result
}

type QueryProcessor interface {
	Process(ctx context.Context, query string) (ProcessorResponse, error)
}

type queryProcessor struct {
	skillRepo        vectorrepository.SkillRepository
	budgetRepo       vectorrepository.BudgetRepository
	experienceRepo   vectorrepository.ExperienceRepository
	miscellanousRepo vectorrepository.MiscellanousRepository
	mercorSkillDb    mysqlrepository.SkillRepository
	openaiProcessor  openai.OpenAiProcessor
}

func NewQueryProcessor() QueryProcessor {
	return queryProcessor{
		skillRepo:        vectorrepository.NewSkillRepository(),
		budgetRepo:       vectorrepository.NewBudgetRepository(),
		experienceRepo:   vectorrepository.NewExperienceRepository(),
		miscellanousRepo: vectorrepository.NewMiscellanousRepository(),
		mercorSkillDb:    mysqlrepository.NewSkillRepository(),
		openaiProcessor:  openai.NewOpenAiProcessor(),
	}
}

func (q queryProcessor) Process(ctx context.Context, query string) (ProcessorResponse, error) {
	embeddingvector, err := q.openaiProcessor.GetEmbeddingVector(query)
	if err != nil {
		return ProcessorResponse{}, err
	}

	budget, err := q.budgetRepo.Get(ctx, embeddingvector)
	if err != nil {
		return ProcessorResponse{}, err
	}

	experience, err := q.experienceRepo.Get(ctx, embeddingvector)
	if err != nil {
		return ProcessorResponse{}, err
	}

	miscellanous, err := q.miscellanousRepo.Get(ctx, embeddingvector)
	if err != nil {
		return ProcessorResponse{}, err
	}

	skills, err := q.skillRepo.Get(ctx, embeddingvector)
	if err != nil {
		return ProcessorResponse{}, err
	}

	budgetRes := ProcessBudget(budget)
	skillRes := ProcessSkills(skills)
	experienceRes := ProcessExperience(experience)
	miscellanousRes := ProcessMiscellanous(miscellanous)
	return ProcessorResponse{
		Skills:       skillRes,
		Budget:       budgetRes,
		Experience:   experienceRes,
		Miscellanous: miscellanousRes,
	}, nil
}
