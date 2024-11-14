package controllers

import (
	services "campaign-service/api/service"
	types "campaign-service/api/types"
	"campaign-service/internal/db"
	"campaign-service/pkg/convert"
	"campaign-service/util"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignController struct {
	service services.CampaignService
}

func NewCampaignController(service services.CampaignService) *CampaignController {
	return &CampaignController{service}
}

// @Param branch body types.BranchesInfoRequest true "Branch information"

// @Summary Create Campaign
// @Tags Campaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Param campaign body types.CampaignRequest true "campaign information"
// @Success 200 {object} util.EmptyResponse "Campaign created successfully"
// @Failure 400 {object} util.EmptyResponse "Invalid input or data"
// @Failure 500 {object} util.EmptyResponse "Failed to create campaign"
// @Router /campaigns [post]
func (c *CampaignController) CreateCampaign(ctx *gin.Context) {
	var campaignRequest types.CampaignRequest
	if err := ctx.ShouldBindJSON(&campaignRequest); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	startDate, _ := convert.ParseDateString("2006-01-02", campaignRequest.StartDate)
	endDate, _ := convert.ParseDateString("2006-01-02", campaignRequest.EndDate)

	campaign := db.Campaign{
		Name:            campaignRequest.Name,
		Description:     campaignRequest.Name,
		StartDate:       startDate,
		EndDate:         endDate,
		MaxParticipants: campaignRequest.MaxParticipants,
	}

	log.Println("data: ", campaign)

	if err := c.service.CreateCampaign(campaign); err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to create campaign")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", nil, "Campaign Created successfully")
}

// @Summary Retrieve a campaign by ID
// @Description This endpoint retrieves a campaign by its ID.
// @Tags Campaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Param id path string true "Campaign ID"
// @Success 200 {object} db.Campaign "Campaign retrieved successfully"
// @Failure 400 {object} util.EmptyResponse "Invalid input or ID format"
// @Failure 500 {object} util.EmptyResponse "Error retrieving Campaign"
// @Router /campaigns/{id} [get]
func (c *CampaignController) GetCampaignDetails(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, "invalid id")
		return
	}
	campaign, err := c.service.GetCampaignDetails(id)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to get campaign by id")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", campaign, "Campaign retrieved successfully")
}

// @Summary Retrieve all campaign
// @Description This endpoint retrieves all campaign.
// @Tags Campaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Success 200 {object} db.Campaign "Campaign retrieved successfully"
// @Failure 500 {object} util.EmptyResponse "Error retrieving Campaign"
// @Router /campaigns [get]
func (c *CampaignController) GetAllCampaignDetails(ctx *gin.Context) {
	campaign, err := c.service.GetAllCampaignDetails()
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to get campaign")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", campaign, "Campaign retrieved successfully")
}
