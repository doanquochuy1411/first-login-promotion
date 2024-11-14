package controllers

import (
	services "campaign-service/api/service"
	types "campaign-service/api/types"
	"campaign-service/internal/db"
	"campaign-service/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignUserController struct {
	service services.CampaignUserService
}

func NewCampaignUserController(service services.CampaignUserService) *CampaignUserController {
	return &CampaignUserController{service}
}

// @Summary Register Campaign
// @Tags RegisterCampaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Param registerInfo body types.CampaignRegistrationInfo true "campaign register information"
// @Success 200 {object} db.Voucher "Campaign registered successfully"
// @Failure 400 {object} util.EmptyResponse "Invalid input or data"
// @Failure 500 {object} util.EmptyResponse "Failed to create campaign"
// @Router /campaigns/register [post]
func (c *CampaignUserController) RegisterCampaign(ctx *gin.Context) {
	var req types.CampaignRegistrationInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, "invalid json")
		return
	}

	campaignRegister := db.CampaignUser{
		Name:               req.Name,
		Email:              req.Email,
		PhoneNumber:        req.PhoneNumber,
		RegistrationMethod: req.RegistrationMethod,
		Status:             "pending",
		CampaignID:         req.CampaignID,
	}

	voucher, err := c.service.CreateCampaignUser(campaignRegister)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to register campaign")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", voucher, "Campaign registered successfully")
}

// @Summary Retrieve a campaign User by ID
// @Description This endpoint retrieves a campaignUser by its ID.
// @Tags RegisterCampaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Param id path string true "CampaignUser ID"
// @Success 200 {object} db.CampaignUser "CampaignUser retrieved successfully"
// @Failure 400 {object} util.EmptyResponse "Invalid input or ID format"
// @Failure 500 {object} util.EmptyResponse "Error retrieving CampaignUser"
// @Router /campaigns/register/{id} [get]
func (c *CampaignUserController) GetCampaignUserDetails(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.RespondWithError(ctx, http.StatusBadRequest, "invalid id")
		return
	}
	campaignUser, err := c.service.GetCampaignUserDetails(id)
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to get campaign user by id")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", campaignUser, "Campaign user retrieved successfully")
}

// @Summary Retrieve All campaign user
// @Description This endpoint retrieves All campaignUser.
// @Tags RegisterCampaign
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param x-api-key header string true "API key for authentication"
// @Success 200 {object} db.CampaignUser "CampaignUser retrieved successfully"
// @Failure 500 {object} util.EmptyResponse "Error retrieving CampaignUser"
// @Router /campaigns/register [get]
func (c *CampaignUserController) GetAllCampaignUserDetails(ctx *gin.Context) {
	campaignUser, err := c.service.GetAllCampaignUser()
	if err != nil {
		util.RespondWithError(ctx, http.StatusInternalServerError, "failed to get campaignUser")
		return
	}

	util.RespondWithJSON(ctx, http.StatusOK, "success", campaignUser, "CampaignUser retrieved successfully")
}
