package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type EmbeddingJSON struct {
	InputID   string    `json:"inputID"`
	InputData string    `json:"inputData"`
	Embedding []float32 `json:"embedding"`
}

type SkillRepository interface {
	Get(ctx context.Context, embedding []float32) ([]Response, error)
	Add(ctx context.Context, data EmbeddingJSON) error
}

type skillRepository struct {
	db *sql.DB
}

func NewSkillRepository() SkillRepository {
	return &skillRepository{
		db: vectordatabase.Initialize(),
	}
}

func (r skillRepository) Add(ctx context.Context, data EmbeddingJSON) error {
	query := fmt.Sprintf("INSERT INTO skills (text, vector) VALUES ('%s', JSON_ARRAY_PACK('[%s]'))", data.InputID, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data.Embedding)), ","), "[]"))
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r skillRepository) Get(ctx context.Context, embedding []float32) ([]Response, error) {
	query := fmt.Sprintf("select text,dot_product(vector, JSON_ARRAY_PACK('[%s]')) as score from skills limit 3 order by score desc", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(embedding)), ","), "[]"))
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	skills := make([]Response, 0)
	// Iterate over the rows and populate the skills slice
	for rows.Next() {
		var skill Response
		if err := rows.Scan(&skill.Text, &skill.Score); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return skills, nil

}
