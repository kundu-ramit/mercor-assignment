package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type ExperienceRepository interface {
	Add(ctx context.Context, data EmbeddingJSON) error
	Get(ctx context.Context, embedding []float32) ([]Response, error)
}

type experienceRepository struct {
	db *sql.DB
}

func NewExperienceRepository() ExperienceRepository {
	return &experienceRepository{
		db: vectordatabase.Initialize(),
	}
}

func (r experienceRepository) Add(ctx context.Context, data EmbeddingJSON) error {
	query := fmt.Sprintf("INSERT INTO experiences (text, vector) VALUES ('%s', JSON_ARRAY_PACK('[%s]'))", data.InputID, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Embedding)), ","), "[]"))
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r experienceRepository) Get(ctx context.Context, embedding []float32) ([]Response, error) {
	query := fmt.Sprintf("select text,dot_product(vector, JSON_ARRAY_PACK('[%s]')) as score from experiences limit 3 order by score desc", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(embedding)), ","), "[]"))
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	experiences := make([]Response, 0)
	// Iterate over the rows and populate the experiences slice
	for rows.Next() {
		var experience Response
		if err := rows.Scan(&experience.Text, &experience.Score); err != nil {
			return nil, err
		}
		experiences = append(experiences, experience)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return experiences, nil

}
