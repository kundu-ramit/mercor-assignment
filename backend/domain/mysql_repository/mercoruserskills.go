package mysqlrepository

import (
	"context"

	"github.com/kundu-ramit/mercor_assignment/infra/database"
	"gorm.io/gorm"
)

type UserSkillMatch struct {
	UserID   string   `json:"userId"`
	SkillIDs []string `json:"skillIds"`
}

type MercorUserSkillRepository interface {
	FetchUserSkillMatches(ctx context.Context, skillIds []string) ([]UserSkillMatch, error)
}

type mercorUserSkillRepository struct {
	db *gorm.DB
}

func NewMercorUserSkillRepository() MercorUserSkillRepository {
	return mercorUserSkillRepository{
		db: database.Initialize(),
	}
}

func (r mercorUserSkillRepository) FetchUserSkillMatches(ctx context.Context, skillIds []string) ([]UserSkillMatch, error) {
	var userSkillMatches []UserSkillMatch

	// Query the database to find distinct userIds with at least one matching skillId
	rows, err := r.db.Table("MercorUserSkills").
		Select("userId, skillId").
		Where("skillId IN (?)", skillIds).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map to store skillIds for each userId
	skillMap := make(map[string][]string)

	// Iterate over the rows and populate skillMap
	for rows.Next() {
		var userID, skillID string
		if err := rows.Scan(&userID, &skillID); err != nil {
			return nil, err
		}
		skillMap[userID] = append(skillMap[userID], skillID)
	}

	// Construct UserSkillMatch objects
	for userID, skillIDs := range skillMap {
		userSkillMatch := UserSkillMatch{
			UserID:   userID,
			SkillIDs: skillIDs,
		}
		userSkillMatches = append(userSkillMatches, userSkillMatch)
	}

	return userSkillMatches, nil
}
