package mysqlrepository

import (
	"context"

	"github.com/kundu-ramit/mercor_assignment/infra/database"
	"gorm.io/gorm"
)

type Skill struct {
	SkillID    string `json:"skillId"`
	SkillName  string `json:"skillName"`
	SkillValue string `json:"skillValue"`
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
	query := "SELECT skillId, skillName, skillValue FROM Skills"

	// Execute raw SQL query with context
	rows, err := r.db.WithContext(ctx).Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var skill Skill
		if err := rows.Scan(&skill.SkillID, &skill.SkillName, &skill.SkillValue); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return skills, nil
}
