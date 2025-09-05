package main

import (
	"fmt"
	"golang-api/core"
	"golang-api/database"
	"golang-api/dotenv"
	"golang-api/user"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MainService struct {
	*core.Provider
	databaseService *database.DatabaseService
	dotenvService   *dotenv.DotenvService
	userController  *user.UserController
}

func NewMainService(
	dbService *database.DatabaseService,
	dt *dotenv.DotenvService,
	userController *user.UserController,
) *MainService {
	return &MainService{
		Provider:        core.NewProvider("MainService"),
		databaseService: dbService,
		dotenvService:   dt,
		userController:  userController,
	}
}

func (ms *MainService) Start() {
	defer ms.databaseService.Close()

	ginMode := ms.dotenvService.Get("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	allowedOrigins := ms.dotenvService.Get("ALLOWED_ORIGINS")

	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(allowedOrigins, ",")
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	config.AllowWebSockets = true
	config.AllowWildcard = true
	config.MaxAge = 0

	r.Use(cors.New(config))

	api := r.Group("/api")
	ms.userController.RegisterRoutes(api)

	if err := r.Run(":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
