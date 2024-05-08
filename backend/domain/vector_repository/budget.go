package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type BudgetRepository interface {
	//Fetch(ctx context.Context, text string) error
	Add(ctx context.Context, data EmbeddingJSON) error
}

type budgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository() BudgetRepository {
	return &budgetRepository{
		db: vectordatabase.Initialize(),
	}
}

func (r budgetRepository) Add(ctx context.Context, data EmbeddingJSON) error {
	query := fmt.Sprintf("INSERT INTO budgets (text, vector) VALUES ('%s', JSON_ARRAY_PACK('[%s]'))", data.InputID, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Embedding)), ","), "[]"))
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
