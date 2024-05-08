package vectorrepository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	vectordatabase "github.com/kundu-ramit/mercor_assignment/infra/vector_database"
)

type MiscellanousRepository interface {
	//Fetch(ctx context.Context, text string) error
	Add(ctx context.Context, data EmbeddingJSON) error
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
