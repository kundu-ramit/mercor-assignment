package scraper

import (
	"context"

	"github.com/kundu-ramit/mercor_assignmentgnmentgnment/domain/entity"
)

type Parser interface {
	Parse(ctx context.Context, html string) (*entity.BullDozer, error)
}
