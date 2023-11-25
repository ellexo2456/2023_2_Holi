package usecase

import (
	"2023_2_Holi/domain"
)

type surveyUsecase struct {
	surveyRepo domain.SurveyRepository
}

func NewSurveyUsecase(sr domain.SurveyRepository) domain.SurveyUsecase {
	return &surveyUsecase{
		surveyRepo: sr,
	}
}

func (u *surveyUsecase) AddSurvey(survey domain.Survey) error {

	if err := u.surveyRepo.AddSurvey(survey); err != nil {
		return err
	} else {
		return nil
	}

}

func (u *surveyUsecase) CheckSurvey(survey domain.Survey) (bool, error) {
	if exist, err := u.surveyRepo.SurveyExists(survey); err != nil {
		return false, err
	} else {
		return exist, err
	}
}
