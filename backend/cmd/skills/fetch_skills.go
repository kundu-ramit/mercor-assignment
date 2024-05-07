package skills

import (
	"encoding/json"
	"log"
	"os"

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
	var skills []Skill
	if err := db.Table("Skills").Select("skillId, skillName").Find(&skills).Error; err != nil {
		panic(err)
	}

	// Create an array of map for skills
	var skillList []map[string]string
	for _, skill := range skills {
		skillMap := map[string]string{
			"skillId":    skill.SkillID,
			"skillName":  skill.SkillName,
			"customText": `Person is skilled in ` + skill.SkillName,
		}
		skillList = append(skillList, skillMap)
	}

	// Marshal the skill list to JSON
	skillJSON, err := json.Marshal(skillList)
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
