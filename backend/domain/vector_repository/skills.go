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
	//Fetch(ctx context.Context, text string) error
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
