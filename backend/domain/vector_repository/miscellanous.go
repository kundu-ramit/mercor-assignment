package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type MiscellanousRepository interface {
	Add(ctx context.Context, data EmbeddingJSON) error
	Get(ctx context.Context, embedding []float32) ([]Response, error)
}

type miscellanousRepository struct {
	db *sql.DB
}

func NewMiscellanousRepository() MiscellanousRepository {
	return &miscellanousRepository{
		db: vectordatabase.Initialize(),
	}
}

func (r miscellanousRepository) Add(ctx context.Context, data EmbeddingJSON) error {
	query := fmt.Sprintf("INSERT INTO miscellanous (text, vector) VALUES ('%s', JSON_ARRAY_PACK('[%s]'))", data.InputID, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Embedding)), ","), "[]"))
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r miscellanousRepository) Get(ctx context.Context, embedding []float32) ([]Response, error) {
	query := fmt.Sprintf("select text,dot_product(vector, JSON_ARRAY_PACK('[%s]')) as score from miscellanous limit 3 order by score desc", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(embedding)), ","), "[]"))
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	miscellanouss := make([]Response, 0)
	// Iterate over the rows and populate the miscellanouss slice
	for rows.Next() {
		var miscellanous Response
		if err := rows.Scan(&miscellanous.Text, &miscellanous.Score); err != nil {
			return nil, err
		}
		miscellanouss = append(miscellanouss, miscellanous)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miscellanouss, nil

}
