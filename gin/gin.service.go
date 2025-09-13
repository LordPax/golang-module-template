package gin

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinService struct {
	*core.Provider
	dotenvService *dotenv.DotenvService
	r             *gin.Engine
	Group         *gin.RouterGroup
}

func NewGinService(module *GinModule) *GinService {
	return &GinService{
		Provider:      core.NewProvider("GinService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (gs *GinService) OnInit() error {
	gs.InitEngine()
	return nil
}

func (gs *GinService) Cors() gin.HandlerFunc {
	allowedOrigins := gs.dotenvService.Get("ALLOWED_ORIGINS")

	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(allowedOrigins, ",")
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	config.AllowWebSockets = true
	config.AllowWildcard = true
	config.MaxAge = 0

	return cors.New(config)
}

func (gs *GinService) InitEngine() {
	fmt.Println("Create gin engine")
	gs.r = gin.Default()

	ginMode := gs.dotenvService.Get("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	gs.r.Use(gin.Logger())
	gs.r.Use(gin.Recovery())
	gs.r.Use(gs.Cors())

	gs.Group = gs.r.Group("/api")
}

func (gs *GinService) Run() {
	port := gs.dotenvService.Get("PORT")
	if err := gs.r.Run(port); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
