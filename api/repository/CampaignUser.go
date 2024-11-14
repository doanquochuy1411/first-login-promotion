package repositories

import (
	"campaign-service/internal/db"
	"campaign-service/pkg/generate"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type CampaignUserRepository interface {
	CreateCampaignUser(campaign db.CampaignUser) (db.Voucher, error)
	GetCampaignUserByID(id int) (db.CampaignUser, error)
	GetAllCampaignUser() ([]db.CampaignUser, error)
}

type campaignUserRepositoryImpl struct {
	store *gorm.DB
}

func NewCampaignUserRepository(store *gorm.DB) CampaignUserRepository {
	return &campaignUserRepositoryImpl{store: store}
}

func (r *campaignUserRepositoryImpl) CreateCampaignUser(campaign db.CampaignUser) (db.Voucher, error) {

	tx := r.store.Begin()
	if tx.Error != nil {
		return db.Voucher{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var result db.CampaignSlots

	err := tx.Raw(`
	SELECT c.ID, c.Name, c.MaxParticipants - COUNT(cu.ID) AS AvailableSlots
	FROM campaigns c
	LEFT JOIN campaign_users cu ON c.ID = cu.CampaignID
	WHERE c.ID = ?
	GROUP BY c.ID, c.Name, c.MaxParticipants;`, campaign.CampaignID).Scan(&result).Error

	if err != nil {
		tx.Rollback()
		return db.Voucher{}, err
	}

	if result.AvailableSlots < 0 {
		tx.Rollback()
		return db.Voucher{}, errors.New("no available slots for the campaign")
	}

	if err := tx.Create(&campaign).Error; err != nil {
		log.Println("Error creating campaign: ", err)
		tx.Rollback()
		return db.Voucher{}, err
	}

	voucher, err := CreateVoucher(30, "available")
	if err != nil {
		log.Println("Error creating voucher: ", err)
		tx.Rollback()
		return db.Voucher{}, err
	}

	if err := tx.Create(&voucher).Error; err != nil {
		log.Println("Error creating voucher: ", err)
		tx.Rollback()
		return db.Voucher{}, err
	}

	tx.Commit()
	return voucher, nil
}

func (r *campaignUserRepositoryImpl) GetCampaignUserByID(id int) (db.CampaignUser, error) {
	var campaign db.CampaignUser
	if err := r.store.First(&campaign, id).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *campaignUserRepositoryImpl) GetAllCampaignUser() ([]db.CampaignUser, error) {
	var campaign []db.CampaignUser
	if err := r.store.Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

// CreateVoucher creates a voucher with a random code and an expiry date 3 days from now
func CreateVoucher(value float64, status string) (db.Voucher, error) {
	code, err := generate.GenerateRandomCode(6)
	if err != nil {
		return db.Voucher{}, fmt.Errorf("failed to generate code: %v", err)
	}

	expiryDate := time.Now().AddDate(0, 0, 3)

	voucher := db.Voucher{
		CodePromotion: code,
		ExpiryDate:    expiryDate,
		Value:         value,
		Status:        status,
		CreatedAt:     generate.PtrTime(time.Now()),
	}

	return voucher, nil
}
