package server

import (
	controllers "campaign-service/api/controller"
	repositories "campaign-service/api/repository"
	services "campaign-service/api/service"
	"campaign-service/internal/middleware"
	"campaign-service/util"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "campaign-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	config util.Config
	store  *gorm.DB
	router *gin.Engine
	server *http.Server
}

func NewServer(cfg util.Config, store *gorm.DB) (*Server, error) {
	server := &Server{
		config: cfg,
		store:  store,
	}
	err := server.setupRouter()
	if err != nil {
		return nil, err
	}

	server.server = &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: server.router,
	}

	return server, nil
}

func (s *Server) setupRouter() error {
	s.router = gin.Default()
	s.router.Use(middleware.CORSMiddleware())

	campaignRepo := repositories.NewCampaignRepository(s.store)
	campaignService := services.NewCampaignService(campaignRepo)
	campaignController := controllers.NewCampaignController(campaignService)

	campaignUserRepo := repositories.NewCampaignUserRepository(s.store)
	campaignUserService := services.NewCampaignUserService(campaignUserRepo)
	campaignUserController := controllers.NewCampaignUserController(campaignUserService)

	private := s.router.Group("/")
	private.Use(middleware.APIKeyMiddleware(s.config.ApiKey))

	private.POST("/campaigns", campaignController.CreateCampaign)
	private.GET("/campaigns", campaignController.GetAllCampaignDetails)
	private.GET("/campaigns/:id", campaignController.GetCampaignDetails)

	private.POST("/campaigns/register", campaignUserController.RegisterCampaign)
	private.GET("/campaigns/register", campaignUserController.GetAllCampaignUserDetails)
	private.GET("/campaigns/register/:id", campaignUserController.GetCampaignUserDetails)

	s.router.GET("/docs", func(c *gin.Context) {
		c.Redirect(302, "/docs/index.html")
	})
	s.router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return nil
}

func (s *Server) Run() error {
	return s.router.Run(s.config.ServerAddress)
}

func (s *Server) Shutdown(sig os.Signal) {
	log.Println("Shutting down server due to signal", logrus.Fields{"signal": sig})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown failed: %+v", err)
	}
	log.Println("Server gracefully stopped")
}
