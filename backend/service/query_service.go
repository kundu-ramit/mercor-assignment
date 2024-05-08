package service

import (
	"context"

	queryprocessor "github.com/kundu-ramit/mercor_assignment/domain/query_processor"
)

type QueryService interface {
	Fetch(ctx context.Context, text string) (queryprocessor.ProcessorResponse, error)
}

type queryService struct {
	queryProcessor queryprocessor.QueryProcessor
}

func NewQueryService() QueryService {
	return queryService{
		queryProcessor: queryprocessor.NewQueryProcessor(),
	}
}

func (q queryService) Fetch(ctx context.Context, text string) (queryprocessor.ProcessorResponse, error) {
	res, err := q.queryProcessor.Process(ctx, text)
	if err != nil {
		return queryprocessor.ProcessorResponse{}, err
	}
	return res, nil
}
