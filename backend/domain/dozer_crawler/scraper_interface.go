package scraper

import (
	"context"

	"github.com/kundu-ramit/mercor_assignment/domain/entity"
)

type Scraper interface {
	ScrapePage(ctx context.Context) ([]entity.BullDozer, error)
}
