package services

import (
	repositories "campaign-service/api/repository"
	"campaign-service/internal/db"
	"errors"
)

type CampaignService interface {
	CreateCampaign(campaign db.Campaign) error
	GetCampaignDetails(id int) (db.Campaign, error)
	GetAllCampaignDetails() ([]db.Campaign, error)
}

type campaignServiceImpl struct {
	repo repositories.CampaignRepository
}

func NewCampaignService(repo repositories.CampaignRepository) CampaignService {
	return &campaignServiceImpl{repo}
}

func (s *campaignServiceImpl) CreateCampaign(campaign db.Campaign) error {
	if campaign.MaxParticipants <= 0 {
		return errors.New("maximum participants must be greater than 0")
	}
	return s.repo.CreateCampaign(campaign)
}

func (s *campaignServiceImpl) GetCampaignDetails(id int) (db.Campaign, error) {
	return s.repo.GetCampaignByID(id)
}

func (s *campaignServiceImpl) GetAllCampaignDetails() ([]db.Campaign, error) {
	return s.repo.GetAllCampaign()
}
