package skills

import (
	"context"
	"encoding/json"
	"log"
	"os"

	mysqlrepository "github.com/kundu-ramit/mercor_assignment/domain/mysql_repository"
	"gorm.io/gorm"
)

// Skill model to map to Skills table
type Skill struct {
	SkillID    string `gorm:"column:skillId"`
	SkillName  string `gorm:"column:skillName"`
	CustomText string `gorm:"column:customText"`
}

func FetchSkills(db *gorm.DB) {

	// Retrieve all records from the Skills table
	skills := make([]Skill, 0)
	dbRes, err := mysqlrepository.NewSkillRepository().FetchAll(context.Background())
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(dbRes); i++ {
		skills = append(skills, Skill{dbRes[i].SkillID, dbRes[i].SkillName, "Person is skilled in " + dbRes[i].SkillName})
	}

	// Marshal the skill list to JSON
	skillJSON, err := json.Marshal(skills)
	if err != nil {
		panic(err)
	}

	// Write JSON to file
	file, err := os.Create("cmd/skills/skill_list.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(skillJSON)
	if err != nil {
		panic(err)
	}

	// Output successful message
	log.Println("Skills list exported to skill_list.json")
}
