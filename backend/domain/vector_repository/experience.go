package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type ExperienceRepository interface {
	//Fetch(ctx context.Context, text string) error
	Add(ctx context.Context, data EmbeddingJSON) error
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
