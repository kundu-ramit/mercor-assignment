package mysqlrepository

import (
	"context"

	"github.com/kundu-ramit/mercor_assignment/infra/database"
	"gorm.io/gorm"
)

type Skill struct {
	SkillID    string
	SkillName  string
	SkillValue string
}

type SkillRepository interface {
	FetchAll(ctx context.Context) ([]Skill, error)
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository() SkillRepository {
	return &skillRepository{
		db: database.Initialize(),
	}
}

func (r *skillRepository) FetchAll(ctx context.Context) ([]Skill, error) {
	var skills []Skill
	query := "SELECT * FROM Skills"
	err := r.db.Raw(query).Scan(&skills).Error
	if err != nil {
		return nil, err
	}
	return skills, nil
}
