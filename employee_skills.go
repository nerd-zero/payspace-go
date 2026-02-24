package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeTraining represents an employee training record.
type EmployeeTraining struct {
	TrainingId     int    `json:"TrainingId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	CourseName     string `json:"CourseName,omitempty"`
	CourseProvider string `json:"CourseProvider,omitempty"`
	StartDate      string `json:"StartDate,omitempty"`
	EndDate        string `json:"EndDate,omitempty"`
	Status         string `json:"Status,omitempty"`
	Result         string `json:"Result,omitempty"`
	ExpiryDate     string `json:"ExpiryDate,omitempty"`
	Cost           string `json:"Cost,omitempty"`
}

// EmployeeQualification represents an employee qualification record.
type EmployeeQualification struct {
	QualificationId   int    `json:"QualificationId"`
	EmployeeNumber    string `json:"EmployeeNumber"`
	QualificationType string `json:"QualificationType,omitempty"`
	Description       string `json:"Description,omitempty"`
	Institution       string `json:"Institution,omitempty"`
	DateObtained      string `json:"DateObtained,omitempty"`
	ExpiryDate        string `json:"ExpiryDate,omitempty"`
}

// EmployeeSkill represents an employee skill record.
type EmployeeSkill struct {
	SkillId        int    `json:"SkillId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	SkillName      string `json:"SkillName,omitempty"`
	SkillLevel     string `json:"SkillLevel,omitempty"`
	YearsExperience string `json:"YearsExperience,omitempty"`
	DateAssessed   string `json:"DateAssessed,omitempty"`
}

// EmployeeSkillsService handles employee skills, training, and qualifications API operations.
type EmployeeSkillsService struct {
	client *Client
}

// --- Training methods (v1.0 for CRUD, v2.0 for ListAll) ---

// ListTraining retrieves all employee training records using v1.0.
func (s *EmployeeSkillsService) ListTraining(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeTraining, *Response, error) {
	s.client.logger.Info("listing employee training", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list training query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURLv1(companyID, "EmployeeTraining")
	var result ListResult[EmployeeTraining]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetTraining retrieves a single training record by ID using v1.0.
func (s *EmployeeSkillsService) GetTraining(ctx context.Context, companyID, trainingID int) (*EmployeeTraining, *Response, error) {
	s.client.logger.Info("getting employee training", zap.Int("company_id", companyID), zap.Int("training_id", trainingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURLv1(companyID, "EmployeeTraining"), trainingID)
	var training EmployeeTraining
	resp, err := s.client.getAndDecode(ctx, url, nil, &training)
	if err != nil {
		return nil, resp, err
	}
	return &training, resp, nil
}

// ListAllTraining retrieves all training records including inactive ones using v2.0.
func (s *EmployeeSkillsService) ListAllTraining(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeTraining, *Response, error) {
	s.client.logger.Info("listing all employee training", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all training query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeTraining/all")
	var result ListResult[EmployeeTraining]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateTraining creates a new training record using v1.0.
func (s *EmployeeSkillsService) CreateTraining(ctx context.Context, companyID int, training *EmployeeTraining) (*EmployeeTraining, *Response, error) {
	s.client.logger.Info("creating employee training", zap.Int("company_id", companyID))

	url := s.client.odataURLv1(companyID, "EmployeeTraining")
	var created EmployeeTraining
	resp, err := s.client.postAndDecode(ctx, url, training, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateTraining updates an existing training record using v1.0.
func (s *EmployeeSkillsService) UpdateTraining(ctx context.Context, companyID, trainingID int, training *EmployeeTraining) (*Response, error) {
	s.client.logger.Info("updating employee training", zap.Int("company_id", companyID), zap.Int("training_id", trainingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURLv1(companyID, "EmployeeTraining"), trainingID)
	return s.client.patch(ctx, url, training)
}

// DeleteTraining deletes a training record using v1.0.
func (s *EmployeeSkillsService) DeleteTraining(ctx context.Context, companyID, trainingID int) (*Response, error) {
	s.client.logger.Info("deleting employee training", zap.Int("company_id", companyID), zap.Int("training_id", trainingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURLv1(companyID, "EmployeeTraining"), trainingID)
	return s.client.del(ctx, url)
}

// --- Qualification methods ---

// ListQualifications retrieves all employee qualifications.
func (s *EmployeeSkillsService) ListQualifications(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeQualification, *Response, error) {
	s.client.logger.Info("listing employee qualifications", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list qualifications query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeQualification")
	var result ListResult[EmployeeQualification]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetQualification retrieves a single qualification by ID.
func (s *EmployeeSkillsService) GetQualification(ctx context.Context, companyID, qualificationID int) (*EmployeeQualification, *Response, error) {
	s.client.logger.Info("getting employee qualification", zap.Int("company_id", companyID), zap.Int("qualification_id", qualificationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeQualification"), qualificationID)
	var qualification EmployeeQualification
	resp, err := s.client.getAndDecode(ctx, url, nil, &qualification)
	if err != nil {
		return nil, resp, err
	}
	return &qualification, resp, nil
}

// ListAllQualifications retrieves all qualifications including inactive ones.
func (s *EmployeeSkillsService) ListAllQualifications(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeQualification, *Response, error) {
	s.client.logger.Info("listing all employee qualifications", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all qualifications query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeQualification/all")
	var result ListResult[EmployeeQualification]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateQualification creates a new qualification record.
func (s *EmployeeSkillsService) CreateQualification(ctx context.Context, companyID int, qualification *EmployeeQualification) (*EmployeeQualification, *Response, error) {
	s.client.logger.Info("creating employee qualification", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeQualification")
	var created EmployeeQualification
	resp, err := s.client.postAndDecode(ctx, url, qualification, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateQualification updates an existing qualification record.
func (s *EmployeeSkillsService) UpdateQualification(ctx context.Context, companyID, qualificationID int, qualification *EmployeeQualification) (*Response, error) {
	s.client.logger.Info("updating employee qualification", zap.Int("company_id", companyID), zap.Int("qualification_id", qualificationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeQualification"), qualificationID)
	return s.client.patch(ctx, url, qualification)
}

// DeleteQualification deletes a qualification record.
func (s *EmployeeSkillsService) DeleteQualification(ctx context.Context, companyID, qualificationID int) (*Response, error) {
	s.client.logger.Info("deleting employee qualification", zap.Int("company_id", companyID), zap.Int("qualification_id", qualificationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeQualification"), qualificationID)
	return s.client.del(ctx, url)
}

// --- Skill methods ---

// ListSkills retrieves all employee skills.
func (s *EmployeeSkillsService) ListSkills(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeSkill, *Response, error) {
	s.client.logger.Info("listing employee skills", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list skills query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeSkill")
	var result ListResult[EmployeeSkill]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetSkill retrieves a single skill by ID.
func (s *EmployeeSkillsService) GetSkill(ctx context.Context, companyID, skillID int) (*EmployeeSkill, *Response, error) {
	s.client.logger.Info("getting employee skill", zap.Int("company_id", companyID), zap.Int("skill_id", skillID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSkill"), skillID)
	var skill EmployeeSkill
	resp, err := s.client.getAndDecode(ctx, url, nil, &skill)
	if err != nil {
		return nil, resp, err
	}
	return &skill, resp, nil
}

// ListAllSkills retrieves all skills including inactive ones.
func (s *EmployeeSkillsService) ListAllSkills(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeSkill, *Response, error) {
	s.client.logger.Info("listing all employee skills", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all skills query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeSkill/all")
	var result ListResult[EmployeeSkill]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateSkill creates a new skill record.
func (s *EmployeeSkillsService) CreateSkill(ctx context.Context, companyID int, skill *EmployeeSkill) (*EmployeeSkill, *Response, error) {
	s.client.logger.Info("creating employee skill", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeSkill")
	var created EmployeeSkill
	resp, err := s.client.postAndDecode(ctx, url, skill, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateSkill updates an existing skill record.
func (s *EmployeeSkillsService) UpdateSkill(ctx context.Context, companyID, skillID int, skill *EmployeeSkill) (*Response, error) {
	s.client.logger.Info("updating employee skill", zap.Int("company_id", companyID), zap.Int("skill_id", skillID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSkill"), skillID)
	return s.client.patch(ctx, url, skill)
}

// DeleteSkill deletes a skill record.
func (s *EmployeeSkillsService) DeleteSkill(ctx context.Context, companyID, skillID int) (*Response, error) {
	s.client.logger.Info("deleting employee skill", zap.Int("company_id", companyID), zap.Int("skill_id", skillID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSkill"), skillID)
	return s.client.del(ctx, url)
}
