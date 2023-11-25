package usecase

import (
	"2023_2_Holi/domain"
	logs "2023_2_Holi/logger"
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

func (u *surveyUsecase) GetStat() ([]domain.Stat, error) {

	stats, err := u.surveyRepo.GetStat()
	if err != nil {
		return nil, err
	}
	logs.Logger.Debug("Usecase GetStat:", stats)

	return stats, nil
}
