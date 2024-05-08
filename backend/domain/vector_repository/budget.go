package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type Response struct {
	Text  string
	Score float32
}

type BudgetRepository interface {
	Add(ctx context.Context, data EmbeddingJSON) error
	Get(ctx context.Context, embedding []float32) ([]Response, error)
}

type budgetRepository struct {
	db *sql.DB
}

func NewBudgetRepository() BudgetRepository {
	return budgetRepository{
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

func (r budgetRepository) Get(ctx context.Context, embedding []float32) ([]Response, error) {
	query := fmt.Sprintf("SELECT text, DOT_PRODUCT(vector, JSON_ARRAY_PACK('[%s]')) AS score FROM budgets  ORDER BY score DESC LIMIT 3", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(embedding)), ","), "[]"))

	fmt.Println(query)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	budgets := make([]Response, 0)
	// Iterate over the rows and populate the budgets slice
	for rows.Next() {
		var budget Response
		if err := rows.Scan(&budget.Text, &budget.Score); err != nil {
			return nil, err
		}
		budgets = append(budgets, budget)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return budgets, nil

}
