package service

import (
	"context"

	mysqlrepository "github.com/kundu-ramit/mercor_assignment/domain/mysql_repository"
	queryprocessor "github.com/kundu-ramit/mercor_assignment/domain/query_processor"
)

type QueryService interface {
	Fetch(ctx context.Context, text string) (queryprocessor.ProcessorResponse, error)
	FetchOrdered(ctx context.Context, skills []string, budget int, experience string) ([]mysqlrepository.UserSkillMatch, error)
}

type queryService struct {
	queryProcessor       queryprocessor.QueryProcessor
	repo                 mysqlrepository.UserDataRepository
	mercorUserSkillsRepo mysqlrepository.MercorUserSkillRepository
}

func NewQueryService() QueryService {
	return queryService{
		queryProcessor:       queryprocessor.NewQueryProcessor(),
		repo:                 mysqlrepository.NewUserDataRepository(),
		mercorUserSkillsRepo: mysqlrepository.NewMercorUserSkillRepository(),
	}
}

func (q queryService) Fetch(ctx context.Context, text string) (queryprocessor.ProcessorResponse, error) {
	res, err := q.queryProcessor.Process(ctx, text)
	if err != nil {
		return queryprocessor.ProcessorResponse{}, err
	}
	return res, nil
}
func (q queryService) FetchOrdered(ctx context.Context, skills []string, budget int, experience string) ([]mysqlrepository.UserSkillMatch, error) {
	userIds, err := q.mercorUserSkillsRepo.FetchUserSkillMatches(ctx, skills)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
