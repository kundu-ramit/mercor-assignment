package mysqlrepository

import (
	"context"

	"github.com/kundu-ramit/mercor_assignment/infra/database"
	"gorm.io/gorm"
)

type UserSkillData struct {
	UserID                 string   `json:"userId"`
	TotalWorkExperience    int      `json:"totalWorkExperience"`
	CompaniesWorkedAt      []string `json:"companiesWorkedAt"`
	OCRGithubUsername      string   `json:"ocrGithubUsername"`
	Schools                []string `json:"schools"`
	Email                  string   `json:"email"`
	Name                   string   `json:"name"`
	Phone                  string   `json:"phone"`
	FullTimeStatus         string   `json:"fullTimeStatus"`
	WorkAvailability       string   `json:"workAvailability"`
	FullTimeSalaryCurrency string   `json:"fullTimeSalaryCurrency"`
	FullTimeSalary         string   `json:"fullTimeSalary"`
	PartTimeSalaryCurrency string   `json:"partTimeSalaryCurrency"`
	PartTimeSalary         string   `json:"partTimeSalary"`
	FullTime               bool     `json:"fullTime"`
	FullTimeAvailability   int      `json:"fullTimeAvailability"`
	PartTime               bool     `json:"partTime"`
	PartTimeAvailability   int      `json:"partTimeAvailability"`
	Skills                 []Skill  `json:"skills"`
	SkillIds               []string
	TotalWorkExperienceDiv string `json:"totalworkexperiencediv"`
}

type MercorUser struct {
	UserID                 string `gorm:"column:userId"`
	Email                  string
	Name                   string
	Phone                  string
	FullTimeStatus         string `gorm:"column:fullTimeStatus"`
	WorkAvailability       string `gorm:"column:workAvailability"`
	FullTimeSalaryCurrency string `gorm:"column:fullTimeSalaryCurrency"`
	FullTimeSalary         string `gorm:"column:fullTimeSalary"`
	PartTimeSalaryCurrency string `gorm:"column:partTimeSalaryCurrency"`
	PartTimeSalary         string `gorm:"column:partTimeSalary"`
	FullTime               bool   `gorm:"column:fullTime"`
	FullTimeAvailability   int    `gorm:"column:fullTimeAvailability"`
	PartTime               bool   `gorm:"column:partTime"`
	PartTimeAvailability   int    `gorm:"column:partTimeAvailability"`
}

type UserResume struct {
	UserID            string `gorm:"column:userId"`
	OCRGithubUsername string `gorm:"column:ocrGithubUsername"`
	ResumeID          string `gorm:"column:resumeId"`
}

type WorkExperience struct {
	ResumeID string `gorm:"column:resumeId"`
	Company  string
}

type Education struct {
	ResumeID string `gorm:"column:resumeId"`
	School   string
}

type UserDataResponse struct {
	UserID   string                 `json:"userId"`
	UserData map[string]interface{} `json:"userData"`
}

type UserDataRepository interface {
	GetUserData(ctx context.Context, userIDs []string, userDataRequests []UserSkillMatch) ([]UserSkillData, error)
}

type userDataRepository struct {
	db *gorm.DB
}

func NewUserDataRepository() UserDataRepository {
	return userDataRepository{
		db: database.Initialize(),
	}
}

func (u userDataRepository) GetUserData(ctx context.Context, userIDs []string, userDataRequests []UserSkillMatch) ([]UserSkillData, error) {
	// Fetch MercorUser and UserResume data in a single query
	var mercorUsers []MercorUser
	var userResumes []UserResume
	if err := u.db.Table("MercorUsers").Where("userId IN (?)", userIDs).Find(&mercorUsers).Error; err != nil {
		return nil, err
	}
	if err := u.db.Table("UserResume").Where("userId IN (?)", userIDs).Find(&userResumes).Error; err != nil {
		return nil, err
	}

	// Fetch WorkExperience and Education data in a single query
	var workExperiences []WorkExperience
	var educations []Education
	if err := u.db.Table("WorkExperience").Where("resumeId IN (?)", getUserResumeIDs(userResumes)).Find(&workExperiences).Error; err != nil {
		return nil, err
	}
	if err := u.db.Table("Education").Where("resumeId IN (?)", getUserResumeIDs(userResumes)).Find(&educations).Error; err != nil {
		return nil, err
	}

	// Construct user data in desired format
	var userSkillData []UserSkillData
	for _, mercorUser := range mercorUsers {
		resumeID := getUserResumeID(mercorUser.UserID, userResumes)

		userData := UserSkillData{
			UserID:                 mercorUser.UserID,
			TotalWorkExperience:    getTotalWorkExperience(resumeID, workExperiences),
			CompaniesWorkedAt:      getCompaniesWorkedAt(resumeID, workExperiences),
			OCRGithubUsername:      getOCRGithubUsername(mercorUser.UserID, userResumes),
			Schools:                getSchools(resumeID, educations),
			Email:                  mercorUser.Email,
			Name:                   mercorUser.Name,
			Phone:                  mercorUser.Phone,
			FullTimeStatus:         mercorUser.FullTimeStatus,
			WorkAvailability:       mercorUser.WorkAvailability,
			FullTimeSalaryCurrency: mercorUser.FullTimeSalaryCurrency,
			FullTimeSalary:         mercorUser.FullTimeSalary,
			PartTimeSalaryCurrency: mercorUser.PartTimeSalaryCurrency,
			PartTimeSalary:         mercorUser.PartTimeSalary,
			FullTime:               mercorUser.FullTime,
			FullTimeAvailability:   mercorUser.FullTimeAvailability,
			PartTime:               mercorUser.PartTime,
			PartTimeAvailability:   mercorUser.PartTimeAvailability,
		}

		// Process the user data requests and add skillIds
		for _, request := range userDataRequests {
			if request.UserID == mercorUser.UserID {
				userData.SkillIds = request.SkillIDs
				break
			}
		}

		userSkillData = append(userSkillData, userData)
	}

	return getUserSkillDataWithReplacements(ctx, userSkillData)
}

// Helper functions

func getUserResumeIDs(userResumes []UserResume) []string {
	ids := make([]string, len(userResumes))
	for i, resume := range userResumes {
		ids[i] = resume.ResumeID
	}
	return ids
}

func getUserResumeID(userID string, userResumes []UserResume) string {
	for _, resume := range userResumes {
		if resume.UserID == userID {
			return resume.ResumeID
		}
	}
	return ""
}

func getTotalWorkExperience(resumeID string, workExperiences []WorkExperience) int {
	count := 0
	for _, experience := range workExperiences {
		if experience.ResumeID == resumeID {
			count++ //cutting corners
		}
	}
	return count
}

func getCompaniesWorkedAt(resumeID string, workExperiences []WorkExperience) []string {
	companies := make(map[string]bool)
	for _, experience := range workExperiences {
		if experience.ResumeID == resumeID {
			companies[experience.Company] = true
		}
	}
	result := make([]string, 0, len(companies))
	for company := range companies {
		result = append(result, company)
	}
	return result
}

func getOCRGithubUsername(userID string, userResumes []UserResume) string {
	for _, resume := range userResumes {
		if resume.UserID == userID {
			return resume.OCRGithubUsername
		}
	}
	return ""
}

func getSchools(resumeID string, educations []Education) []string {
	schools := make(map[string]bool)
	for _, edu := range educations {
		if edu.ResumeID == resumeID {
			schools[edu.School] = true
		}
	}
	result := make([]string, 0, len(schools))
	for school := range schools {
		result = append(result, school)
	}
	return result
}

func getUserSkillDataWithReplacements(ctx context.Context, data []UserSkillData) ([]UserSkillData, error) {

	skillIds, _ := NewSkillRepository().FetchAll(ctx)

	m := make(map[string]Skill)
	for i := range skillIds {
		m[skillIds[i].SkillID] = skillIds[i]
	}

	for i := range data {
		// Replace skill IDs with skills
		for j, skillId := range data[i].SkillIds {
			data[i].Skills[j] = m[skillId]
		}
		// Replace experience level ID with description
		data[i].TotalWorkExperienceDiv = getExperienceDescription(data[i].TotalWorkExperience)
	}
	return data, nil
}

func getExperienceDescription(experience int) string {
	if experience <= 2 {
		return "BEG"
	}
	if experience <= 8 {
		return "INT"
	}
	return "EXP"
}
