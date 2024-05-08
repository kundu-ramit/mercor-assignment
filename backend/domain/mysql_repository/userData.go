package mysqlrepository

import (
	"github.com/kundu-ramit/mercor_assignment/infra/database"
	"gorm.io/gorm"
)

type UserDataRequest struct {
	UserID   string   `json:"userId"`
	SkillIds []string `json:"skillIds"`
}

type UserDataResponse struct {
	UserID   string                 `json:"userId"`
	UserData map[string]interface{} `json:"userData"`
}

type UserDataRepository interface {
	GetUserData(userDataRequests []UserDataRequest) ([]UserDataResponse, error)
}

type userDataRepository struct {
	db *gorm.DB
}

func NewUserDataRepository() UserDataRepository {
	return userDataRepository{
		db: database.Initialize(),
	}
}

func (u userDataRepository) GetUserData(userDataRequests []UserDataRequest) ([]UserDataResponse, error) {
	var userDataResponses []UserDataResponse

	for _, request := range userDataRequests {
		// Fetch MercorUser and UserResume data in a single query
		var mercorUser MercorUser
		var userResume UserResume
		if err := u.db.Table("MercorUsers").Where("userId = ?", request.UserID).Find(&mercorUser).Error; err != nil {
			return nil, err
		}
		if err := u.db.Table("UserResume").Where("userId = ?", request.UserID).Find(&userResume).Error; err != nil {
			return nil, err
		}

		// Fetch WorkExperience and Education data in a single query
		var workExperiences []WorkExperience
		var educations []Education
		if err := u.db.Table("WorkExperience").Where("resumeId = ?", userResume.ResumeID).Find(&workExperiences).Error; err != nil {
			return nil, err
		}
		if err := u.db.Table("Education").Where("resumeId = ?", userResume.ResumeID).Find(&educations).Error; err != nil {
			return nil, err
		}

		// Construct user data in desired format
		userData := map[string]interface{}{
			"userId":                 mercorUser.UserID,
			"totalWorkExperience":    getTotalWorkExperience(userResume.ResumeID, workExperiences),
			"companiesWorkedAt":      getCompaniesWorkedAt(userResume.ResumeID, workExperiences),
			"ocrGithubUsername":      userResume.OCRGithubUsername,
			"schools":                getSchools(userResume.ResumeID, educations),
			"email":                  mercorUser.Email,
			"name":                   mercorUser.Name,
			"phone":                  mercorUser.Phone,
			"fullTimeStatus":         mercorUser.FullTimeStatus,
			"workAvailability":       mercorUser.WorkAvailability,
			"fullTimeSalaryCurrency": mercorUser.FullTimeSalaryCurrency,
			"fullTimeSalary":         mercorUser.FullTimeSalary,
			"partTimeSalaryCurrency": mercorUser.PartTimeSalaryCurrency,
			"partTimeSalary":         mercorUser.PartTimeSalary,
			"fullTime":               mercorUser.FullTime,
			"fullTimeAvailability":   mercorUser.FullTimeAvailability,
			"partTime":               mercorUser.PartTime,
			"partTimeAvailability":   mercorUser.PartTimeAvailability,
			"skillIds":               request.SkillIds, // Add skillIds to userData
		}

		userDataResponse := UserDataResponse{
			UserID:   mercorUser.UserID,
			UserData: userData,
		}

		userDataResponses = append(userDataResponses, userDataResponse)
	}

	return userDataResponses, nil
}

// Helper functions

func getTotalWorkExperience(resumeId string, workExperiences []WorkExperience) int {
	count := 0
	for _, experience := range workExperiences {
		if experience.ResumeID == resumeId {
			count++
		}
	}
	return count
}

func getCompaniesWorkedAt(resumeId string, workExperiences []WorkExperience) []string {
	companies := make(map[string]bool)
	for _, experience := range workExperiences {
		if experience.ResumeID == resumeId {
			companies[experience.Company] = true
		}
	}
	result := make([]string, 0, len(companies))
	for company := range companies {
		result = append(result, company)
	}
	return result
}

func getSchools(resumeId string, educations []Education) []string {
	schools := make(map[string]bool)
	for _, edu := range educations {
		if edu.ResumeID == resumeId {
			schools[edu.School] = true
		}
	}
	result := make([]string, 0, len(schools))
	for school := range schools {
		result = append(result, school)
	}
	return result
}
