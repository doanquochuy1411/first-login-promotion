package types

type CampaignRequest struct {
	ID              int     `json:"id" example:"ignore"`
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	DiscountRate    float64 `json:"discount_rate" binding:"required"`
	StartDate       string  `json:"start_date" binding:"required"`
	EndDate         string  `json:"end_date" binding:"required"`
	MaxParticipants int     `json:"max_participants" binding:"required"`
}

type CampaignRegistrationInfo struct {
	CampaignID         int    `json:"id" example:"ignore"`
	Name               string `json:"name" binding:"required"`
	Email              string `json:"email" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	RegistrationMethod string `json:"registration_method" binding:"required"`
}
