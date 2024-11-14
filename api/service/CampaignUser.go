package services

import (
	repositories "campaign-service/api/repository"
	"campaign-service/internal/db"
)

type CampaignUserService interface {
	CreateCampaignUser(campaignUser db.CampaignUser) (db.Voucher, error)
	GetCampaignUserDetails(id int) (db.CampaignUser, error)
	GetAllCampaignUser() ([]db.CampaignUser, error)
}

type campaignUserServiceImpl struct {
	repo repositories.CampaignUserRepository
}

func NewCampaignUserService(repo repositories.CampaignUserRepository) CampaignUserService {
	return &campaignUserServiceImpl{repo}
}

func (s *campaignUserServiceImpl) CreateCampaignUser(campaignUser db.CampaignUser) (db.Voucher, error) {
	if campaignUser.RegistrationMethod == "link" {
		campaignUser.SubscriptionType = "sliver"
	} else {
		campaignUser.SubscriptionType = "normal"
	}
	return s.repo.CreateCampaignUser(campaignUser)
}

func (s *campaignUserServiceImpl) GetCampaignUserDetails(id int) (db.CampaignUser, error) {
	return s.repo.GetCampaignUserByID(id)
}

func (s *campaignUserServiceImpl) GetAllCampaignUser() ([]db.CampaignUser, error) {
	return s.repo.GetAllCampaignUser()
}
