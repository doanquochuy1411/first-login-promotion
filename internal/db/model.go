package db

import (
	"time"
)

type Campaign struct {
	ID              int        `gorm:"primaryKey;autoIncrement"`
	Name            string     `gorm:"type:varchar(100);not null"`
	Description     string     `gorm:"type:text"`
	StartDate       time.Time  `gorm:"type:timestamp without time zone;not null"`
	EndDate         time.Time  `gorm:"type:timestamp without time zone;not null"`
	MaxParticipants int        `gorm:"not null"`
	CreatedAt       *time.Time `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"deleted_at"`
}

type CampaignUser struct {
	ID                 int    `gorm:"primary_key;auto_increment"`
	Name               string `gorm:"not null"`
	PhoneNumber        string `gorm:"not null;unique"`
	Email              string `gorm:"not null;unique"`
	RegistrationMethod string `gorm:"type:varchar(20);not null"`
	Status             string `gorm:"type:varchar(50);"`
	SubscriptionType   string `gorm:"type:varchar(50);"`

	CreatedAt *time.Time `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"deleted_at"`

	CampaignID int       `gorm:"index;not null"`
	Campaign   *Campaign `gorm:"foreignKey:CampaignID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// type SubscriptionType struct {
// 	ID   int    `gorm:"primary_key;auto_increment"`
// 	Name string `gorm:"not null"`

// 	VoucherID int      `gorm:"null"`
// 	Voucher   *Voucher `gorm:"foreignKey:VoucherID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
// }

// type User struct {
// 	ID          int        `gorm:"primary_key;auto_increment"`
// 	Name        string     `gorm:"not null"`
// 	PhoneNumber string     `gorm:"not null"`
// 	Email       string     `gorm:"not null"`
// 	CreatedAt   *time.Time `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt   *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"updated_at"`
// 	DeletedAt   *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"deleted_at"`
// }

type Voucher struct {
	ID            int        `gorm:"primary_key;auto_increment"`
	CodePromotion string     `gorm:"type:varchar(50);unique"`
	ExpiryDate    time.Time  `gorm:"not null"`
	Value         float64    `gorm:"not null"`
	Status        string     `gorm:"type:varchar(30);not null"`
	CreatedAt     *time.Time `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"type:timestamp without time zone;autoUpdateTime" json:"updated_at"`
}

// Custom type
type CampaignSlots struct {
	ID             int
	Name           string
	AvailableSlots int
}
