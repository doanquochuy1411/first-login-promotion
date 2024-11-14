package repositories

import (
	"campaign-service/internal/db"
	"log"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	CreateCampaign(campaign db.Campaign) error
	GetCampaignByID(id int) (db.Campaign, error)
	GetAllCampaign() ([]db.Campaign, error)
}

type campaignRepositoryImpl struct {
	store *gorm.DB
}

func NewCampaignRepository(store *gorm.DB) CampaignRepository {
	return &campaignRepositoryImpl{store: store}
}

func (r *campaignRepositoryImpl) CreateCampaign(campaign db.Campaign) error {
	err := r.store.Create(&campaign).Error
	if err != nil {
		log.Println("Error creating campaign: ", err)
		return err
	}
	return nil
}

func (r *campaignRepositoryImpl) GetCampaignByID(id int) (db.Campaign, error) {
	var campaign db.Campaign
	if err := r.store.First(&campaign, id).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *campaignRepositoryImpl) GetAllCampaign() ([]db.Campaign, error) {
	var campaign []db.Campaign
	if err := r.store.Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}
